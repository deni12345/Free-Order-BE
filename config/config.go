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
	os.Setenv("FIREBASE_CREDENTIAL", "ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAiZnJlZS1vcmRlci1mYiIsCiAgInByaXZhdGVfa2V5X2lkIjogImFiYTEyMzVmN2Q1MGFlZWZkZjM4NmVmNGMwOTA1ZWNmMmU3ZTBlZDUiLAogICJwcml2YXRlX2tleSI6ICItLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRRHhMUkEzZ3VsUExkUFlcbkJ6ZzdsdFRWMi9JcW1OUEFiMWJHUjRPR3FHdll0S0JMNURqbXU3Rk1GYTRjWE4vZzFldFR5Y2ZXY05GRWtkVTVcbnFGc1gwdUw5aVZGeGFBTjdsMmFmUjcyNU5WSDZCdE5JL1phS2ZVSC8zZWJIQWpjamg1R1FrckZNKzc3Qjl5N0lcbm5FSnQzTnBLN2I5ZGxncGNxLzJZbFV1VGpkNlRlVTN1cktCRVlMSUFibDNZbndlZHBhV1p6Y1BhWEljd1crRlJcbjB3YU5xZERVMEsrYjRtVjhOcVZWRnRMOVB0ZlBQbS9VZkZkWGt6YnZ2a3JXZVRLMjNhS1drcGVuTzRGaDRtR1RcbkFxaEppS2krTmNHN3JpMGx2TkNZVThNUjBqK2xGbjhVbXpPaTBCRDZzMEs3d1JPVUZ1QnE1MTRzVzJXYWlUNCtcblVacUlieGNSQWdNQkFBRUNnZ0VBRkJiUDJmL2lsVURZKzQ2bVpYaEw1RUI4N0lWc1pnUWdNMFh6czd6RnR2aGNcbnZkV1YwOHBnbTFqVzl1L1ppaTNuZTBROGpDVTdtVDJZaUVQNHlvUTdlaWo5Q3JoQ1NnSXpmeU1IcjZ5OFZscEpcbk1UUkdYT1RRSVpNSFRSNW5IMk9FZ1lsbWtXWUszcmRMUTFTWXgvUDBqdmM4MEh2V0F5QWtjSmhIQWxSMzdNcXRcbmlEZTZ3cGhsb3RTazREdjl6RUZQb3NLc1QrbUJCQWlCTmR4ekEvVk4zN2xoRWw4b1o0QkFyNmQxVEVXMXMxclhcbmE3VzRHOUQ5UnBaZUpTTTBieUZJL05Ua3J5V3dVNkZ4SWJhMVBCUGVzOHl4K2RwUUtVQkh2QkNPMWJZdGxxRkZcbjcxdnh2ZXduMVdwcGl5aVhZeHlma0w5Yk5XY3B5YmNHM3c3c1pucnBYUUtCZ1FEN2YzKzAzU0k0MnB1WFU5aTlcblB0clRaS0NPOTNyd1lFTlpkSXFNVHBzWEdiK0U1K1A3c2FrdlZSUGl2NmlyK0N5RXVFSkpETmhvRjlkNnNMQzFcbjZ5UzR2ZHZsQnFJTkJyRWYwaXkrTzUzYWRPWnFDbXdMMVBHckRyQ2dYUHozeTNFcmpIbWlnVzVZcHVJZWpDNnRcbndUZmJTdEVzR0ZKTzhOakVaSzI4NFgwakl3S0JnUUQxZmtOdUlhVDVsU09XNFFYNDh1MjNWV0VGcnJJV2U1MExcbmIwaUh1NlNFSU4xRGViRklDdlNmQ1o1QnpiM1lnUzRFNEczV0lpa3h6ODQwcXJiVUFUNGpBWkNZUlRyOFZqdDlcbnVZY0tDLzFXOURsRCtJWjBueUpyUVIreGlUaGJGZzJZS0F5Qm5VaUFmdHJsdmVSWVBQSi9JUHM3cC9qT2wwU3ZcbkJtODZaRDdxT3dLQmdRRGd0VjZwZnh0Z3FSSno5MXFrWEZDd3FDdzlFU29yTkJkNUNnUHdzUnNwWEx3ZVNBTVBcbmxBR1NaUFhMbDJ0aG0vQks5VWRrZGJHMTcvZDdpL3VYejIyVUV5OHlSdWJydEpyMXo1aGlzR211RXR6OHlpc1hcbnR6L21rczJGdXcvYlowN0tsa3pud2JXU1NDbXBHWjlyYUVROEQzRXRjTjI1NHNBTFdkbFI5amNVK1FLQmdFUURcbmlGRUIxMEFpbHFxZGkxY09qdGVsT3UwbEdrZnFWWGN1akt5MlN2MGtVY2o0OVNuZGh5cHNzc0RPYkpPSUxxVEJcblBRei9oK1E2QzRwQjVjZFUxTTlFQnJoNUxiOXRjS3dEYzZ0UDZROW5PRnBoaDNiV0ErWXRNOGV2R2NMNm1DZVlcbnFPWmNHaE1ua2lQbWVWWU5CSzQrQi96ZUs2dW84eWwrb2VCT2w3T3pBb0dBVVRlLzZIUHNMQU9VMno0YjZhS2NcbkVjMUFxemF4WEJnc0pRZk12VkJMUlhuNkdaeFZRVnNwQWVBdnNJTWVMb2MvRDNFSVRXSmRycVdLSFBGMXh2MGJcbkRtbEkxUTc0N1pJbmNFOEcxaTZ0djQ0ZHNzeURBQmZqazlqOHppaGtqcFc3TDVzVDgvaVdTZzVwUVpRM3ZVdjJcbjRPaGNtSjJkeXpDZDl3OVU0Rk05c1U0PVxuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwKICAiY2xpZW50X2VtYWlsIjogImZpcmViYXNlLWFkbWluc2RrLWZqMnBpQGZyZWUtb3JkZXItZmIuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJjbGllbnRfaWQiOiAiMTA4NDkzMjM5MTMxMzUzMjI3MjY2IiwKICAiYXV0aF91cmkiOiAiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tL28vb2F1dGgyL2F1dGgiLAogICJ0b2tlbl91cmkiOiAiaHR0cHM6Ly9vYXV0aDIuZ29vZ2xlYXBpcy5jb20vdG9rZW4iLAogICJhdXRoX3Byb3ZpZGVyX3g1MDlfY2VydF91cmwiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vb2F1dGgyL3YxL2NlcnRzIiwKICAiY2xpZW50X3g1MDlfY2VydF91cmwiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vcm9ib3QvdjEvbWV0YWRhdGEveDUwOS9maXJlYmFzZS1hZG1pbnNkay1majJwaSU0MGZyZWUtb3JkZXItZmIuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJ1bml2ZXJzZV9kb21haW4iOiAiZ29vZ2xlYXBpcy5jb20iCn0K")
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
