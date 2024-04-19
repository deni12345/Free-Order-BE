SAMPLE_BINARY_NAME=main

MIGRATE_SOURCE=file://database/script
DATABASE='mysql://root:Liemdjack1@@tcp(localhost:3306)/fodb?multiStatements=true'

build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./cmd/main.go
	$(USERPROFILE)\Go\bin\build-lambda-zip.exe -o main.zip main
zip: 
	$(USERPROFILE)\Go\bin\build-lambda-zip.exe -o main.zip main
run:
	go run ./cmd/$(SAMPLE_BINARY_NAME).go
migrate:
	migrate -source $(MIGRATE_SOURCE) -database $(DATABASE) up 