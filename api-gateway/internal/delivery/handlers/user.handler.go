package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/savanyv/digital-wallet/proto/user"
	"google.golang.org/grpc"
)

var userClient pb.UserServiceClient

func init() {
	conn, err := grpc.Dial("user-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	userClient = pb.NewUserServiceClient(conn)
}

func GetProfileUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := userClient.GetUser(ctx, &pb.GetUserRequest{
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
