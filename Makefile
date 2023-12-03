build:
	go build -o bin/TIKTRADER

run:
	go build -o bin/TIKTRADER

	./bin/TIKTRADER

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto

.PHONY: proto