package main

import (
	"fmt"
	"net"

	"github.com/brenddonanjos/clean_commerce/services/payment/configs"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/infra/database"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/infra/grpc/pb"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/infra/grpc/service"
	"github.com/brenddonanjos/clean_commerce/services/payment/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Connect(
		config.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName),
	)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	cardRepository := database.NewCardRepository(db)
	createCardUsecase := usecase.NewCreateCardUseCase(cardRepository)

	grpcServer := grpc.NewServer()
	createCardService := service.NewCardService(createCardUsecase)
	pb.RegisterCardServiceServer(grpcServer, createCardService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port 50051...")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(listener)
}
