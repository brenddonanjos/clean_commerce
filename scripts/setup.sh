# Generate on services
protoc ---go_out=. --go-grpc_out=. -proto_path=internal/infra/grpc --grpc-gateway_out=. internal/infra/grpc/proto/user.proto

#Generateon gateway-service
protoc --go_out=. --go-grpc_out=. --grpc-gateway_opt=paths=source_relative --proto_path=../payment/internal/infra/grpc/proto --proto_path=internal/infra/grpc --grpc-gateway_out=./internal/infra/grpc/pb_payment card.proto

protoc --go_out=. --go-grpc_out=. --grpc-gateway_opt=paths=source_relative --proto_path=../user/internal/infra/grpc/proto --proto_path=internal/infra/grpc --grpc-gateway_out=./internal/infra/grpc/pb_user user.proto

protoc --go_out=. --go-grpc_out=. --grpc-gateway_opt=paths=source_relative --proto_path=../user/internal/infra/grpc/proto --proto_path=internal/infra/grpc --grpc-gateway_out=./internal/infra/grpc/pb_user address.proto



# libprotoc 3.12.4
# protoc-gen-go v1.34.2
# protoc-gen-go-grpc 1.5.1


