package configs

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	RootPath = filepath.Join(filepath.Dir(b), "..")
)

type Config struct {
	// App related config
	ServerToRun string `mapstructure:"SERVER_TO_RUN"`
	AppName     string `mapstructure:"APP_NAME"`
	AppHttpPort string `mapstructure:"APP_HTTP_PORT"`
	AppGrpcPort string `mapstructure:"APP_GRPC_PORT"`
	// DB related config
	DbHost string `mapstructure:"DB_HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbName string `mapstructure:"DB_NAME"`
	DbUser string `mapstructure:"DB_USER"`
	DbPass string `mapstructure:"DB_PASS"`
}

func LoadConfig() (config Config, err error) {
	// tell viper the location of the config file
	viper.SetConfigFile(filepath.Join(RootPath, "configs", ".env"))

	// tell viper the type of the config file, which is env
	viper.SetConfigType("env")

	// override values that it has red from config file with values of the corresponding environment
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
