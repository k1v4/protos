all: sso

sso:
	@echo "Generating gRPC and grpc-gateway"
	protoc -I ./proto \
      --go_out ./gen/ --go_opt paths=source_relative \
      --go-grpc_out ./gen/ --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./gen/ --grpc-gateway_opt paths=source_relative \
      ./proto/file_uploader/uploader.proto

