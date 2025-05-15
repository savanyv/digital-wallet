package grpcdelivery

import (
	"context"

	dtos "github.com/savanyv/digital-wallet/auth-service/internal/dto"
	"github.com/savanyv/digital-wallet/auth-service/internal/usecase"
	pb "github.com/savanyv/digital-wallet/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	usecase usecase.AuthUsecase
}

func NewAuthServer(usecase usecase.AuthUsecase) *AuthServer {
	return &AuthServer{
		usecase: usecase,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	result, err := s.usecase.Register(&dtos.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.AuthResponse{
		UserId:   result.UserId,
		Message:  result.Message,
	}

	return resp, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	result, err := s.usecase.Login(&dtos.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.AuthResponse{
		UserId:   result.UserId,
		Token:    result.Token,
		Message:  result.Message,
	}

	return resp, nil
}
