package handlers

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/savanyv/digital-wallet/proto/transaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var transactionClient pb.TransactionServiceClient

func init() {
	conn, err := grpc.Dial("transaction-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	transactionClient = pb.NewTransactionServiceClient(conn)
}

func Deposit(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	reqBody := new(pb.DepositWithdrawRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := transactionClient.Deposit(ctx, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}

func Withdraw(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	reqBody := new(pb.DepositWithdrawRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := transactionClient.Withdraw(ctx, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}

func Transfer(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		authHeader = "Bearer " + strings.TrimSpace(authHeader)
	}

	md := metadata.New(map[string]string{
		"authorization": authHeader,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reqBody := new(pb.TransferRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := transactionClient.Transfer(ctx, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}

func GetTransactionHistory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := transactionClient.GetTransactionHistory(ctx, &pb.GetHistoryRequest{
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
