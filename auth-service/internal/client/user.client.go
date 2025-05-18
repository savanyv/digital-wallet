package client

import (
	"context"
	"log"
	"time"

	userPB "github.com/savanyv/digital-wallet/proto/user"

	"google.golang.org/grpc"
)

type UserGrpcClient interface {
	CreateUser(ctx context.Context, req *userPB.CreateUserRequest) (*userPB.UserResponse, error)
}

type userGrpcClient struct {
	client userPB.UserServiceClient
}

func NewUserGrpcClient() (UserGrpcClient, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	client := userPB.NewUserServiceClient(conn)
	return &userGrpcClient{
		client: client,
	}, nil
}

func (c *userGrpcClient) CreateUser(ctx context.Context, req *userPB.CreateUserRequest) (*userPB.UserResponse, error) {
	resp, err := c.client.CreateUser(ctx, req)
	if err != nil {
		log.Printf("Failed to create user via gRPC: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *userGrpcClient) FindUserByEmail(ctx context.Context, req *userPB.GetUserByEmailRequest) (*userPB.UserResponse, error) {
	resp, err := c.client.GetUserByEmail(ctx, req)
	if err != nil {
		log.Printf("Failed to get user via gRPC: %v", err)
		return nil, err
	}

	return resp, nil
}
