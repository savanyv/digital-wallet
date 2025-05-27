package client

import (
	"context"
	"log"
	"time"

	walletPB "github.com/savanyv/digital-wallet/proto/wallet"
	"google.golang.org/grpc"
)

type WalletGrpcClient interface {
	CreateWallet(ctx context.Context, req *walletPB.CreateWalletRequest) (*walletPB.WalletResponse, error)
}

type walletGrpcClient struct {
	client walletPB.WalletServiceClient
}

func NewWalletGrpcClient() (WalletGrpcClient, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	client := walletPB.NewWalletServiceClient(conn)
	return &walletGrpcClient{
		client: client,
	}, nil
}

func (c *walletGrpcClient) CreateWallet(ctx context.Context, req *walletPB.CreateWalletRequest) (*walletPB.WalletResponse, error) {
	resp, err := c.client.CreateWallet(ctx, req)
	if err != nil {
		log.Printf("Failed to create wallet via gRPC: %v", err)
		return nil, err
	}

	return resp, nil
}
