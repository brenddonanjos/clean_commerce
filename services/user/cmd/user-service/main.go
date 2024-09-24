package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/brenddonanjos/clean_commerce/services/user/configs"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/database"
	pb "github.com/brenddonanjos/clean_commerce/services/user/internal/infra/grpc/pb_user"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/grpc/service"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open(
		config.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	repository := database.New(db)

	//user
	createUserUseCase := usecase.NewCreateUserUseCase(repository)
	createUserService := service.NewUserService(createUserUseCase)
	pb.RegisterUserServiceServer(grpcServer, createUserService)

	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC User service on port 50052...")
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(listener)

}
