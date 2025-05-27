package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/digital-wallet/api-gateway/internal/delivery/handlers"
	"github.com/savanyv/digital-wallet/shared/middlewares"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
)

func walletServiceRoutes(app fiber.Router) {
	jwtService := jwt.NewJWTService()

	app.Get("/wallets/:user_id", middlewares.AuthMiddlewares(jwtService), handlers.GetWallet)
}
