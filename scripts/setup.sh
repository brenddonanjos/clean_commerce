protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/card.proto

protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. internal/infra/grpc/proto/card.proto

protoc -I . -I googleapis --go_out=. --go-grpc_out=. --grpc-gateway_out=. services/payment/internal/infra/grpc/proto/card.proto



protoc -I services/payment/internal/infra/grpc/proto \
       -I services/googleapis \
       --go_out=services/payment/internal/infra/grpc/proto \
       --go-grpc_out=services/payment/internal/infra/grpc/proto \
       --grpc-gateway_out=services/payment/internal/infra/grpc/proto \
       services/payment/internal/infra/grpc/proto/card.proto


protoc --proto_path=. --proto_path=./googleapis --go_out=. --go-grpc_out=. --grpc-gateway_out=. services/payment/internal/infra/grpc/proto/card.proto

protoc --proto_path=internal/infra/grpc --go_out=. --go-grpc_out=. --grpc-gateway_out=. internal/infra/grpc/proto/card.proto

