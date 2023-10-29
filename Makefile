SAMPLE_BINARY_NAME=main

build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./cmd/main.go
	$(USERPROFILE)\Go\bin\build-lambda-zip.exe -o main.zip main
zip: 
	$(USERPROFILE)\Go\bin\build-lambda-zip.exe -o main.zip main
run:
	go run ./cmd/$(SAMPLE_BINARY_NAME).go