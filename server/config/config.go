package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PG_URI  string `mapstructure:"PG_URI"`
	PORT    string `mapstructure:"PORT"`
	JWT_KEY string `mapstructure:"JWT_KEY"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			PG_URI:  os.Getenv("PG_URI"),
			PORT:    os.Getenv("PORT"),
			JWT_KEY: os.Getenv("JWT_KEY"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// validate config here
	if config.PG_URI == "" {
		err = errors.New("PG_URI is required")
		return
	}

	if config.PORT == "" {
		err = errors.New("PORT is required")
		return
	}

	if config.JWT_KEY == "" {
		err = errors.New("JWT_KEY is required")
		return
	}

	return
}
