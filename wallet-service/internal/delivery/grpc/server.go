package grpcdelivery

import (
	"context"

	pb "github.com/savanyv/digital-wallet/proto/wallet"
	dtos "github.com/savanyv/digital-wallet/wallet-service/internal/dto"
	"github.com/savanyv/digital-wallet/wallet-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletServer struct {
	pb.UnimplementedWalletServiceServer
	usecase usecase.WalletUsecase
}

func NewWalletServer(usecase usecase.WalletUsecase) *WalletServer {
	return &WalletServer{
		usecase: usecase,
	}
}

func (u *WalletServer) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.WalletResponse, error) {
	result, err := u.usecase.CreateWallet(&dtos.CreateWalletRequest{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.WalletResponse{
		UserId:   result.UserID,
		Balance:  result.Balance,
		Message:  result.Message,
	}

	return resp, nil
}

func (u *WalletServer) GetWallet(ctx context.Context, req *pb.GetWalletRequest) (*pb.WalletResponse, error) {
	result, err := u.usecase.GetWallet(&dtos.GetWalletRequest{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.WalletResponse{
		UserId: result.UserID,
		Balance: result.Balance,
		Message: result.Message,
	}

	return resp, nil
}

func (u *WalletServer) UpdateBalance(ctx context.Context, req *pb.UpdateBalanceRequest) (*pb.WalletResponse, error) {
	result, err := u.usecase.UpdateBalance(&dtos.UpdateBalanceRequest{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.WalletResponse{
		UserId: result.UserID,
		Balance: result.Balance,
		Message: result.Message,
	}

	return resp, nil
}
