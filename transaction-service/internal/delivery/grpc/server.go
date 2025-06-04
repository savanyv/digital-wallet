package grpcdelivery

import (
	"context"
	"time"

	pb "github.com/savanyv/digital-wallet/proto/transaction"
	dtos "github.com/savanyv/digital-wallet/transaction-service/internal/dto"
	"github.com/savanyv/digital-wallet/transaction-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionServer struct {
	pb.UnimplementedTransactionServiceServer
	usecase usecase.TransactionUsecase
}

func NewTransactionServer(usecase usecase.TransactionUsecase) *TransactionServer {
	return &TransactionServer{
		usecase: usecase,
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
