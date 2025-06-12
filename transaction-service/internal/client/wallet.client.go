package client

import (
	"context"
	"log"
	"time"

	walletPb "github.com/savanyv/digital-wallet/proto/wallet"
	"google.golang.org/grpc"
)

type WalletGrpcClient interface {
	UpdateBalance(ctx context.Context, req *walletPb.UpdateBalanceRequest) (*walletPb.WalletResponse, error)
}

type walletGrpcClient struct {
	client walletPb.WalletServiceClient
}

func NewWalletGrpcClient() (WalletGrpcClient, error) {
	conn, err := grpc.Dial("wallet-service:50052", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	client := walletPb.NewWalletServiceClient(conn)
	return &walletGrpcClient{
		client: client,
	}, nil
}

func (c *walletGrpcClient) UpdateBalance(ctx context.Context, req *walletPb.UpdateBalanceRequest) (*walletPb.WalletResponse, error) {
	resp, err := c.client.UpdateBalance(ctx, req)
	if err != nil {
		log.Printf("Failed to update balance via gRPC: %v", err)
		return nil, err
	}

	return resp, nil
}
