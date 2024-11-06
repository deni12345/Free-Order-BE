SAMPLE_BINARY_NAME=main

MIGRATE_SOURCE=file://database/script
DATABASE='mysql://root:password@tcp(localhost:3306)/fodb?charset=utf8mb4&parseTime=True&loc=Local'

install:
	go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest
	
build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./cmd/main.go

zip: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main ./cmd/main.go
	$(GOPATH)\bin\build-lambda-zip.exe -o build/main.zip build/main

run:
	go run ./cmd/$(SAMPLE_BINARY_NAME).go

nodemon:
	nodemon --watch './**/*.go' --signal SIGKILL --exec 'go' run cmd/*.go

migrate:
	migrate -source $(MIGRATE_SOURCE) -database $(DATABASE) up

tidy:
	go mod tidy && go mod vendor

lint:
	golangci-lint run