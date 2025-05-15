package app

import (
	"log"
	"net"

	"github.com/savanyv/digital-wallet/auth-service/internal/database"
	grpcdelivery "github.com/savanyv/digital-wallet/auth-service/internal/delivery/grpc"
	"github.com/savanyv/digital-wallet/auth-service/internal/repository"
	"github.com/savanyv/digital-wallet/auth-service/internal/usecase"
	"github.com/savanyv/digital-wallet/shared/config"
	"google.golang.org/grpc"
	pb "github.com/savanyv/digital-wallet/proto/auth"
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
	usecase := usecase.NewAuthUsecase(repo)

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
