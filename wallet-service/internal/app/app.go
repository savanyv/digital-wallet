package app

import (
	"log"
	"net"

	"github.com/savanyv/digital-wallet/shared/config"
	"github.com/savanyv/digital-wallet/wallet-service/internal/database"
	grpcdelivery "github.com/savanyv/digital-wallet/wallet-service/internal/delivery/grpc"
	"github.com/savanyv/digital-wallet/wallet-service/internal/repository"
	"github.com/savanyv/digital-wallet/wallet-service/internal/usecase"
	"google.golang.org/grpc"
	pb "github.com/savanyv/digital-wallet/proto/wallet"
)

func Run() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	walletRepo := repository.NewWalletRepository(db)
	usecase := usecase.NewWalletUsecase(walletRepo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	walletServer := grpcdelivery.NewWalletServer(usecase)

	s := grpc.NewServer()
	pb.RegisterWalletServiceServer(s, walletServer)

	log.Printf("Wallet-Service gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
