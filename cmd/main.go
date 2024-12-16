package main

import (
	"context"
	"github/free-order-be/api"
	"github/free-order-be/config"
	"github/free-order-be/internal/dao"
	"github/free-order-be/tool"
	"net/http"
	"os"

	_ "github/free-order-be/docs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	muxadapter "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/sirupsen/logrus"
)

var (
	daoInst   *dao.DAO
	muxLambda *muxadapter.GorillaMuxAdapterV2
)

func init() {
	var err error
	config.InitLogrus()
	config.LoadConfig()

	conn := config.Values.ConnectDB(context.Background())
	if daoInst, err = dao.NewDAO(context.Background(), conn); err != nil {
		logrus.Errorf("error create dao: %s", err)
	}

	tables, _ := conn.ListTables().All(context.Background())
	logrus.Infof("list table on database %v", tables)
}

func main() {
	server := api.NewServer(daoInst)

	if os.Getenv("RUN_ENV") == config.DEV || os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		logrus.Info("run server in lambda mode [:8080]")
		muxLambda = muxadapter.NewV2(server.Router)
		lambda.Start(Handler)
		return
	}

	tool.ShowBanner()
	logrus.Info("run server in localhost mode [:8080]")
	if err := http.ListenAndServe("localhost:8080", server.Router); err != nil {
		logrus.Fatalf("error starting server: %s", err)
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return muxLambda.ProxyWithContext(ctx, req)
}
