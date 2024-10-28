package config

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/logging"
	"github.com/guregu/dynamo/v2"
	"github.com/sirupsen/logrus"
)

func (cfg *configValue) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (smithyendpoints.Endpoint, error) {
	var (
		endpoint *url.URL
		err      error
	)
	if cfg.Env == LOCAL {
		endpoint, err = url.Parse(cfg.DynamodbEndpoint)
		if err != nil {
			return smithyendpoints.Endpoint{}, err
		}
	} else {
		endpoint, err = url.Parse(fmt.Sprintf("https://stg.dynamodb.%s.amazonaws.com", cfg.DB.Region))
		if err != nil {
			return smithyendpoints.Endpoint{}, err
		}
	}
	return smithyendpoints.Endpoint{URI: *endpoint}, nil
}

func (cfg *configValue) Logf(classification logging.Classification, format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

func (cfg *configValue) ApplyResolveEnpoint(o *dynamodb.Options) {
	o.EndpointResolverV2 = cfg
}

func (cfg *configValue) ApplyLogrusLogger(o *dynamodb.Options, mode aws.ClientLogMode) {
	o.ClientLogMode = mode
	o.Logger = cfg
}

func (cfg *configValue) ConnectDB(ctx context.Context) *dynamo.DB {
	awsConfig, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		logrus.Fatalf("load default aws config on err: %v", err)
	}

	return dynamo.New(awsConfig, func(o *dynamodb.Options) {
		cfg.ApplyResolveEnpoint(o)
		cfg.ApplyLogrusLogger(o, aws.LogRequestWithBody)
	})
}
