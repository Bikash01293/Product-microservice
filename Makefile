proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:.

product-server:
	go run cmd/main.go