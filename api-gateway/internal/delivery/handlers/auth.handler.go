package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/savanyv/digital-wallet/proto/auth"
	"google.golang.org/grpc"
)

var authClient pb.AuthServiceClient

func  init() {
	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	authClient = pb.NewAuthServiceClient(conn)
}

func RegisterAuthHandler(c *fiber.Ctx) error {
	type Req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	reqBody := new(Req)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := authClient.Register(ctx, &pb.RegisterRequest{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": resp,
	})
}

func LoginAuthHandler(c *fiber.Ctx) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	reqBody := new(Req)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := authClient.Login(ctx, &pb.LoginRequest{
		Email:    reqBody.Email,
		Password: reqBody.Password,
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
