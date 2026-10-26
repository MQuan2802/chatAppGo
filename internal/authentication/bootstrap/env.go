package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Port          int    `mapstructure:"PORT"` // Added to catch the server port
	Datasource    string `mapstructure:"DATA_SOURCE"`
	DbPort        int    `mapstructure:"DATA_SOURCE_PORT"`
	DbTimeout     int    `mapstructure:"DB_TIMEOUT"`
	DbUsername    string `mapstructure:"DB_USERNAME"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
	DbMaxAttempts int    `mapstructure:"DB_MAX_ATTEMPTS"`
	DbMaxPoolSize int    `mapstructure:"DB_MAX_POOL_SIZE"`
}

var ENV_PATH = ".env"

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigName(".env")
	// Tell Viper what format it is
	viper.SetConfigType("env")

	// 📂 Look for the file in the project root directory (current working directory)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	readConfigErr := viper.ReadInConfig()

	if readConfigErr != nil {
		log.Fatal("Can not find the config file")
	}

	marshalErr := viper.Unmarshal(&env)

	if marshalErr != nil {
		log.Fatal("Can not load environment")
	}
	if env.AppEnv == "development" {
		log.Printf("App is running in development environment")
	}
	return &env
}
