package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	Value = configValue{
		SecretKey: []byte("nones"),
		DB: struct {
			DBName   string "yaml:\"db_name\""
			UserName string "yaml:\"user_name\""
			Password string "yaml:\"password\""
			Host     string "yaml:\"host\""
			Port     string "yaml:\"port\""
		}{},
	}
)

type configValue struct {
	SecretKey          []byte `yaml:"secret_key"`
	GoogleClientID     string `yaml:"google_client_id"`
	GoogleClientSecret string `yaml:"google_client_secret"`
	DB                 struct {
		DBName   string `yaml:"db_name"`
		UserName string `yaml:"user_name"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	}
}

func LoadConfig(env string) (configValue, error) {
	values := configValue{}
	content, err := ioutil.ReadFile(fmt.Sprintf(`./config/%s.yaml`, env))
	if err != nil {
		return values, fmt.Errorf("error: %v", err)
	}

	err = yaml.Unmarshal(content, &values)
	if err != nil {
		return values, fmt.Errorf("error: %v", err)
	}
	return values, nil
}
