# Generate on payment-service
protoc --proto_path=internal/infra/grpc --go_out=. --go-grpc_out=. --grpc-gateway_out=. internal/infra/grpc/proto/card.proto

#Generateon gateway-service
protoc --go_out=. --go-grpc_out=. --grpc-gateway_opt=paths=source_relative --proto_path=../payment/internal/infra/grpc/proto --proto_path=. --grpc-gateway_out=./internal/infra/grpc/pb card.proto



# libprotoc 3.12.4
# protoc-gen-go v1.34.2
# protoc-gen-go-grpc 1.5.1


