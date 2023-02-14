package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PG_URI       string `mapstructure:"PG_URI"`
	PORT         string `mapstructure:"PORT"`
	JWT_KEY      string `mapstructure:"JWT_KEY"`
	AWS_REGION   string `mapstructure:"AWS_REGION"`
	AWS_ENDPOINT string `mapstructure:"AWS_ENDPOINT"`
	QUEUE        string `mapstructure:"QUEUE"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			PG_URI:       os.Getenv("PG_URI"),
			PORT:         os.Getenv("PORT"),
			JWT_KEY:      os.Getenv("JWT_KEY"),
			AWS_REGION:   os.Getenv("AWS_REGION"),
			AWS_ENDPOINT: os.Getenv("AWS_ENDPOINT"),
			QUEUE:        os.Getenv("QUEUE"),
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

	if config.AWS_ENDPOINT == "" {
		err = errors.New("AWS_ENDPOINT is required")
		return
	}

	if config.AWS_REGION == "" {
		err = errors.New("AWS_REGION is required")
		return
	}

	if config.QUEUE == "" {
		err = errors.New("QUEUE is required")
		return
	}

	return
}
