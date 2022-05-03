package config

import (
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

func NewConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("healthchecker")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error on load configuration file")
	}

	return Config{
		Viper: viper.GetViper(),
	}
}
