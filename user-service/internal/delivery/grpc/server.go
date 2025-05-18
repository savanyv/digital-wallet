package grpcdelivery

import (
	"context"

	pb "github.com/savanyv/digital-wallet/proto/user"
	dtos "github.com/savanyv/digital-wallet/user-service/internal/dto"
	"github.com/savanyv/digital-wallet/user-service/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	usecase usecase.UserUsecase
}

func NewUserServer(usecase usecase.UserUsecase) *UserServer {
	return &UserServer{
		usecase: usecase,
	}
}

func (u *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	result, err := u.usecase.CreateUser(&dtos.CreateUserRequest{
		UserID: req.UserId,
		Name:   req.Name,
		Email:  req.Email,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UserResponse{
		UserId:    result.UserID,
		Name:  result.Name,
		Email: result.Email,
	}

	return resp, nil
}

func (u *UserServer) FindUserByID(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	result, err := u.usecase.FindUserByID(req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UserResponse{
		UserId:    result.UserID,
		Name:  result.Name,
		Email: result.Email,
	}

	return resp, nil
}

func (u *UserServer) FindUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	result, err := u.usecase.FindUserByEmail(req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UserResponse{
		UserId:    result.UserID,
		Name:  result.Name,
		Email: result.Email,
	}

	return resp, nil
}
