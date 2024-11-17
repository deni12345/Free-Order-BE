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
	logrus "github.com/sirupsen/logrus"
)

func (cfg *configValue) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (smithyendpoints.Endpoint, error) {
	var err error
	var endpointURL *url.URL

	// endpointString := cfg.DynamodbEndpoint
	// if cfg.Env != LOCAL && cfg.DB.Region != "" {
	// 	endpointString = fmt.Sprintf("https://dynamodb.%s.amazonaws.com", cfg.DB.Region)
	// }
	endpointString := fmt.Sprintf("https://dynamodb.%s.amazonaws.com", cfg.DB.Region)
	if endpointURL, err = url.Parse(endpointString); err != nil {
		return smithyendpoints.Endpoint{}, err
	}

	return smithyendpoints.Endpoint{URI: *endpointURL}, nil
}

func (cfg *configValue) Logf(classification logging.Classification, format string, v ...interface{}) {
	logrus.Info(v...)
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
