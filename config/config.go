package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	DEV          = "dev"
	LOCAL        = "local"
	UNDEFINE_ENV = ""
)

var (
	Values = &configValue{}
)

type database struct {
	Region string `yaml:"region" envconfig:"REGION"`
}

type configValue struct {
	Env                string
	SecretKey          string   `yaml:"secret_key" envconfig:"SECRET_KEY"`
	GoogleClientID     string   `yaml:"google_client_id" envconfig:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string   `yaml:"google_client_secret" envconfig:"GOOGLE_CLIENT_SECRET"`
	DynamodbEndpoint   string   `yaml:"dynamodb_endpoint" envconfig:"DYNAMODB_ENDPOINT"`
	DB                 database `yaml:"db"`
}

func LoadConfig() {
	Environment := os.Getenv("RUN_ENV")
	if Environment == UNDEFINE_ENV {
		Environment = LOCAL
	}
	Values = loadConfigValues(Environment)

}

func loadConfigValues(env string) *configValue {
	values := &configValue{}
	values.Env = env
	content, err := os.ReadFile(fmt.Sprintf(`./config/%s.yaml`, env))
	if err != nil {
		logrus.Fatalf("error read config yaml file for %s: %v", env, err)
		return nil
	}

	err = envconfig.Process("", values)
	if err != nil {
		logrus.Fatalf("error process envconfig value for %s: %v", env, err)
		return nil
	}

	err = yaml.Unmarshal(content, values)
	if err != nil {
		logrus.Fatalf("error umarshal yaml config file for %s: %v", env, err)
		return nil
	}
	return values
}
