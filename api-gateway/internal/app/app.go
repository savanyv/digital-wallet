package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/digital-wallet/api-gateway/internal/delivery/routes"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	return &Server{
		app: fiber.New(),
	}
}

func (s *Server) Run() error {
	// Routes
	routes.RegisterRoutes(s.app)

	// Start
	if err := s.app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	return nil
}
