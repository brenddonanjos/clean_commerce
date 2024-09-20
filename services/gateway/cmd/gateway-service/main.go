package main

import (
	"context"
	"log"
	"net/http"

	"github.com/brenddonanjos/clean_commerce/services/gateway/internal/infra/grpc/pb_payment"
	"github.com/brenddonanjos/clean_commerce/services/gateway/internal/infra/grpc/pb_user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// var (
// 	paymentEndpoint = flag.String("payment-endpoint", "payment:50051", "gRPC payment endpoint")
// )

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	//Payment service
	err := pb_payment.RegisterCardServiceHandlerFromEndpoint(ctx, mux, "payment:50051", opts)
	if err != nil {
		panic(err)
	}

	//User service
	err = pb_user.RegisterAddressServiceHandlerFromEndpoint(ctx, mux, "user:50052", opts)
	if err != nil {
		panic(err)
	}

	err = pb_user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "user:50052", opts)
	if err != nil {
		panic(err)
	}

	log.Println("Gateway está rodando na porta 8000...")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("Falha ao iniciar o servidor HTTP: %v", err)
	}
}
