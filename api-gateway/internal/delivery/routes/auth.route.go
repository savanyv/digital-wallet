package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/digital-wallet/api-gateway/internal/delivery/handlers"
)

func authServiceRoutes(app fiber.Router) {
	app.Post("/register", handlers.RegisterAuthHandler)
	app.Post("/login", handlers.LoginAuthHandler)
}
