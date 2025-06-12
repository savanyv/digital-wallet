package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/savanyv/digital-wallet/proto/wallet"
	"google.golang.org/grpc"
)

var walletClient pb.WalletServiceClient

func init() {
	conn, err := grpc.Dial("wallet-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	walletClient = pb.NewWalletServiceClient(conn)
}

func GetWallet(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := walletClient.GetWallet(ctx, &pb.GetWalletRequest{
		UserId: c.Params("user_id"),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}
