package config

import (
	"healthchecker/adapters/logger"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
)

type Config struct {
	Viper *viper.Viper
}

func NewConfig(logger logger.Logger) Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("healthchecker")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error on load configuration file")
	}

	logger.Info("Config loaded")

	return Config{
		Viper: viper.GetViper(),
	}
}
