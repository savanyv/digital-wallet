package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/digital-wallet/api-gateway/internal/delivery/handlers"
	"github.com/savanyv/digital-wallet/shared/middlewares"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
)

func transactionServiceRoutes(app fiber.Router) {
	jwtService := jwt.NewJWTService()

	app.Post("/transactions/deposit", middlewares.AuthMiddlewares(jwtService), handlers.Deposit)
	app.Post("/transactions/withdraw", middlewares.AuthMiddlewares(jwtService), handlers.Withdraw)
	app.Post("transactions/transfer", middlewares.AuthMiddlewares(jwtService), handlers.Transfer)
	app.Get("/transactions/:user_id", middlewares.AuthMiddlewares(jwtService), handlers.GetTransactionHistory)
}
