// Protobuf installation instructions https://github.com/golang/protobuf#installation
generate:
	protoc --proto_path=proto  --go_out=plugins=grpc,import_path=auth:proto/auth proto/service.proto

run:
	go run cmd/gpp-grpc-demo/main.go

test:
	go test -v ./...