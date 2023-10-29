package main

import (
	"context"
	"github/lambda-microservice/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var (
	client    *s3.Client
	ginLambda *ginadapter.GinLambda
)

func main() {
	// bucket := os.Getenv("BUCKET_NAME")
	// fileName := os.Getenv("FILE_NAME")
	// region := os.Getenv("REGION")

	// cfg, err := config.LoadDefaultConfig(context.TODO(),
	// 	config.WithRegion(region),
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// client = s3.NewFromConfig(cfg)
	r := buildEngine()
	routes.SetUserRoutes(r)

	ginLambda = ginadapter.New(r)
	lambda.Start(handler)
}

func buildEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return engine
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
