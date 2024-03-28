package config

import (
	"errors"
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres struct {
		DSN string `env:"POSTGRES_DSN"`
	}
	Redis struct {
		Password string `env:"REDIS_PASSWORD"`
		Port     string `env:"REDIS_PORT"`
		DSN      string `env:"REDIS_DSN"`
	}
	Api struct {
		Port        string `env:"API_PORT"`
		Environment string
	}
}

func SelectEnvFile(instance *Config, environment string) (string, error) {
	switch environment {
	case "prod":
		instance.Api.Environment = "prod"
		return ".env", nil
	case "dev":
		instance.Api.Environment = "dev"
		return ".dev.env", nil
	default:
		return "", errors.New("invalid environment specified")
	}
}

func New() (*Config, error) {
	var environment string
	flag.StringVar(&environment, "environment", "prod", "Specify the environment (prod or dev)")
	flag.Parse()

	var instance Config
	envFile, err := SelectEnvFile(&instance, environment)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig(envFile, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
