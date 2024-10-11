SAMPLE_BINARY_NAME=main

MIGRATE_SOURCE=file://database/script
DATABASE='mysql://root:password@tcp(localhost:3306)/fodb?charset=utf8mb4&parseTime=True&loc=Local'

build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./cmd/main.go
#	$(USERPROFILE)\Go\bin\build-lambda-zip.exe -o main.zip main
zip: 
	$(USERPROFILE)\Go\bin\build-lambda-zip.exe -o main.zip main
run:
	go run ./cmd/$(SAMPLE_BINARY_NAME).go
# run watch:
# 	nodemon --watch './**/*.go' --signal SIGKILL --exec 'go' run cmd/*.go
migrate:
	migrate -source $(MIGRATE_SOURCE) -database $(DATABASE) up 
tidy:
	go mod tidy && go mod vendor