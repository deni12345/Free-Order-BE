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
	GoogleID           string   `yaml:"google_id" envconfig:"GOOGLE_ID"`
	GoogleClientSecret string   `yaml:"google_client_secret" envconfig:"GOOGLE_CLIENT_SECRET"`
	DynamodbEndpoint   string   `yaml:"dynamodb_endpoint" envconfig:"DYNAMODB_ENDPOINT"`
	DB                 database `yaml:"db"`
	FirebaseCredential string   `yaml:"firebase_credential" envconfig:"FIREBASE_CREDENTIAL"`
	RedirectURL        string   `yaml:"redirect_url" envconfig:"REDIRECT_URL"`
}

func LoadConfig() {
	Environment := os.Getenv("RUN_ENV")
	if Environment == UNDEFINE_ENV {
		Environment = LOCAL
	}
	loadConfigValues(Environment, Values)
}

func loadConfigValues(env string, values *configValue) *configValue {
	values.Env = env

	content, err := os.ReadFile(fmt.Sprintf(`./config/%s.yaml`, env))
	if err != nil {
		logrus.Infof("error read config yaml file for %s: %v", env, err)
	}

	if err = envconfig.Process("", values); err != nil {
		logrus.Infof("error process envconfig value for %s: %v", env, err)
		return nil
	}
	if err = yaml.Unmarshal(content, values); err != nil {
		logrus.Infof("error umarshal yaml config file for %s: %v", env, err)
		return nil
	}

	return values
}

func InitLogrus() {
	logrus.StandardLogger().SetFormatter(&logrus.TextFormatter{
		EnvironmentOverrideColors: true,
		ForceColors:               true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		PadLevelText:              true,
	})
}
