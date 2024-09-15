package main

import (
	"github/lambda-microservice/api"
	"github/lambda-microservice/custom"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	server := api.NewServer()

	custom.PrintLogo()
	logrus.Info("run server in localhost mode [:8080]")
	http.ListenAndServe("localhost:8080", server.Router)
}
