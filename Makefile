SAMPLE_BINARY_NAME=main

MIGRATE_SOURCE=file://database/script
DATABASE='mysql://root:password@tcp(localhost:3306)/fodb?charset=utf8mb4&parseTime=True&loc=Local'

install:
	go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	go install github.com/swaggo/swag/cmd/swag@latest
	
build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main ./cmd/main.go && cp --parents ./tool/ascii-art.txt ./build

zip: build
	$(GOPATH)\bin\build-lambda-zip.exe -o build/main.zip build/main

lint:
	golangci-lint run --fix

swag: 
	@echo "Formating swagger anotations..."
	@swag fmt
	@echo "Generate swagger file..."
	@swag init -g api.go -d ./*

run:
	go run ./cmd/$(SAMPLE_BINARY_NAME).go

nodemon:
	nodemon --watch './**/*.go' --signal SIGKILL --exec 'go' run cmd/*.go

migrate:
	migrate -source $(MIGRATE_SOURCE) -database $(DATABASE) up

tidy:
	go mod tidy && go mod vendor