package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	paymentPb "github.com/brenddonanjos/clean_commerce/services/payment/internal/infra/grpc/pb"
)

var (
	paymentEndpoint = flag.String("payment-endpoint", "payment:50051", "Api gateway is ON!")
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := paymentPb.RegisterCardServiceHandlerFromEndpoint(ctx, mux, "payment:50051", opts)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8081", mux)
}
