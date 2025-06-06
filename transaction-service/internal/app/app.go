package app

import (
	"log"
	"net"

	pb "github.com/savanyv/digital-wallet/proto/transaction"
	"github.com/savanyv/digital-wallet/shared/config"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
	"github.com/savanyv/digital-wallet/transaction-service/internal/client"
	"github.com/savanyv/digital-wallet/transaction-service/internal/database"
	grpcdelivery "github.com/savanyv/digital-wallet/transaction-service/internal/delivery/grpc"
	"github.com/savanyv/digital-wallet/transaction-service/internal/repository"
	"github.com/savanyv/digital-wallet/transaction-service/internal/usecase"
	"google.golang.org/grpc"
)

func Run() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.NewTransactionRepository(db)
	walletClient, err := client.NewWalletGrpcClient()
	if err != nil {
		log.Fatalf("failed to connect to wallet service: %v", err)
	}
	usecase := usecase.NewTransactionUsecase(repo, walletClient)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	jwtService := jwt.NewJWTService()
	transactionServer := grpcdelivery.NewTransactionServer(usecase, jwtService)

	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, transactionServer)

	log.Printf("Transaction-Service gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
