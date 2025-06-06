package grpcdelivery

import (
	"context"
	"strings"
	"time"

	pb "github.com/savanyv/digital-wallet/proto/transaction"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
	dtos "github.com/savanyv/digital-wallet/transaction-service/internal/dto"
	"github.com/savanyv/digital-wallet/transaction-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TransactionServer struct {
	pb.UnimplementedTransactionServiceServer
	usecase usecase.TransactionUsecase
	jwtService jwt.JWTService
}

func NewTransactionServer(usecase usecase.TransactionUsecase, jwtService jwt.JWTService) *TransactionServer {
	return &TransactionServer{
		usecase: usecase,
		jwtService: jwtService,
	}
}

func (s *TransactionServer) Deposit(ctx context.Context, req *pb.DepositWithdrawRequest) (*pb.TransactionResponse, error) {
	result, err := s.usecase.Deposit(&dtos.DepositWithdrawRequest{
		UserID: req.UserId,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.TransactionResponse{
		UserId: result.UserID,
		Type: result.Type,
		Amount: result.Amount,
		Message: result.Message,
		CreatedAt: result.CreatedAt.Format(time.RFC3339),
	}

	return resp, nil
}

func (s *TransactionServer) Withdraw(ctx context.Context, req *pb.DepositWithdrawRequest) (*pb.TransactionResponse, error) {
	result, err := s.usecase.Withdraw(&dtos.DepositWithdrawRequest{
		UserID: req.UserId,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.TransactionResponse{
		UserId: result.UserID,
		Type: result.Type,
		Amount: int64(result.Amount),
		Message: result.Message,
		CreatedAt: result.CreatedAt.Format(time.RFC3339),
	}

	return resp, nil
}

func (s *TransactionServer) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransactionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
	    return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
	    return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	tokenParts := strings.Split(strings.TrimSpace(authHeader[0]), " ")
	if len(tokenParts) != 2 || !strings.EqualFold(tokenParts[0], "Bearer") {
	    return nil, status.Error(codes.Unauthenticated, "invalid token format")
	}

	token := tokenParts[1]
	if token == "" {
	    return nil, status.Error(codes.Unauthenticated, "empty token")
	}

	claims, err := s.jwtService.ValidateToken(token)
	if err != nil {
	    return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	userID := claims.UserID

	result, err := s.usecase.Transfer(&dtos.TransferRequest{
		SenderID: userID,
		ReceiverID: req.ReceiverId,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.TransactionResponse{
		UserId: result.UserID,
		Type: result.Type,
		Amount: result.Amount,
		Message: result.Message,
		CreatedAt: result.CreatedAt.Format(time.RFC3339),
	}

	return resp, nil
}
