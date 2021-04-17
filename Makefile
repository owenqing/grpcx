
update:
	@echo "start to update proto file..."
	@protoc --go_out=plugins=grpc:./pb ./proto/*.proto
	@echo "update success"

build:
	@echo "start to build project..."
	@go build -o ./bin/grpcx ./main.go
	@echo "end build"


run: 
	@echo "start to run project..."
	@go run ./cmd/main.go

ssl: 
	@echo "gen ssl..."
	@openssl req -x509 -newkey rsa:4096 -keyout ./cert/key.pem -out ./cert/cert.pem -days 365 -nodes -subj /CN=localhost
	@echo "end ssl"
	
