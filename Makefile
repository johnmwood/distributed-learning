.PHONY: run
run:
	go run cmd/server/main.go
	# grpcurl -plaintext -d '{"Key": "something"}' localhost:50051 bora.BoraService/GetValue

.PHONY: test
test:
	go test -v ./...

.PHONY: protos
protos:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/bora.proto
