package main

import (
	"github/free-order-be/api"
	"github/free-order-be/custom"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

func main() {
	server := api.NewServer()
	custom.PrintLogo()

	if os.Getenv("RUN_ENV") == "lambda" || os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		logrus.Info("run server in lambda mode [:8080]")
		lambda.Start(server.Router)
		return
	}

	logrus.Info("run server in localhost mode [:8080]")
	if err := http.ListenAndServe("localhost:8080", server.Router); err != nil {
		logrus.Fatal("error starting server: %s", err)
	}
}
