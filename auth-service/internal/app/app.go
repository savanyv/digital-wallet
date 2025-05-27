package app

import (
	"log"
	"net"

	"github.com/savanyv/digital-wallet/auth-service/internal/client"
	"github.com/savanyv/digital-wallet/auth-service/internal/database"
	grpcdelivery "github.com/savanyv/digital-wallet/auth-service/internal/delivery/grpc"
	"github.com/savanyv/digital-wallet/auth-service/internal/repository"
	"github.com/savanyv/digital-wallet/auth-service/internal/usecase"
	pb "github.com/savanyv/digital-wallet/proto/auth"
	"github.com/savanyv/digital-wallet/shared/config"
	"google.golang.org/grpc"
)

func Run() {
	// Config
	cfg := config.LoadConfig()

	// Database
	db, err := database.ConnectPosgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// initialize repo and usecase
	repo := repository.NewAuthRepository(db)
	userClient, err := client.NewUserGrpcClient()
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	walletClient, err := client.NewWalletGrpcClient()
	if err != nil {
		log.Fatalf("failed to connect to wallet service: %v", err)
	}
	usecase := usecase.NewAuthUsecase(repo, userClient, walletClient)

	// initialize grpc server
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authServer := grpcdelivery.NewAuthServer(usecase)

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, authServer)

	log.Printf("Auth-Service gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
