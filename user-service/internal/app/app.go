package app

import (
	"log"
	"net"

	pb "github.com/savanyv/digital-wallet/proto/user"
	"github.com/savanyv/digital-wallet/shared/config"
	"github.com/savanyv/digital-wallet/user-service/internal/database"
	grpcdelivery "github.com/savanyv/digital-wallet/user-service/internal/delivery/grpc"
	"github.com/savanyv/digital-wallet/user-service/internal/repository"
	"github.com/savanyv/digital-wallet/user-service/internal/usecase"
	"google.golang.org/grpc"
)

func Run() {
	// Load Config
	cfg := config.LoadConfig()

	// Connect to Database
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// initialize repo and usecase
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecsae(repo)

	// initialize grpc server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userServer := grpcdelivery.NewUserServer(usecase)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userServer)

	log.Printf("User-Service gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
