package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	appEnv         string `mapstructure:"APP_ENV"`
	serverAddress  string `mapstructure:"SERVER_ADDRESS"`
	datasource     string `mapstructure:"DATA_SOURCE"`
	dbPort         int    `mapstructure:"DATA_SOURCE_PORT"`
	contextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	dbUsername     string `mapstructure:"DB_USERNAME"`
	dbPassword     string `mapstructure:"DB_STRUCTURE"`
	dbName         string `mapstructure:"DB_NAME"`
}

var ENV_PATH = ".env"

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(ENV_PATH)

	readConfigErr := viper.ReadInConfig()

	if readConfigErr != nil {
		log.Fatal("Can not find the config file")
	}

	marshalErr := viper.Unmarshal()

	if marshalErr != nil {
		log.Fatal("Can not load environment")
	}
	if env.appEnv == "development" {
		log.Printf("App is running in development environment")
	}
	return &env
}
