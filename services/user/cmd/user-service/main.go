package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/database"
	pb "github.com/brenddonanjos/clean_commerce/services/user/internal/infra/grpc/pb_user"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/infra/grpc/service"
	"github.com/brenddonanjos/clean_commerce/services/user/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/commerce_ai")
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
	//address
	createAddressUseCase := usecase.NewCreateAddressUseCase(repository)
	createAddressService := service.NewAddressService(createAddressUseCase)
	pb.RegisterAddressServiceServer(grpcServer, createAddressService)

	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC User service on port 50052...")
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(listener)

}
