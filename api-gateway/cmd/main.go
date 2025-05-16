package main

import (
	"log"

	"github.com/savanyv/digital-wallet/api-gateway/internal/app"
)

func main() {
	server := app.NewServer()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
