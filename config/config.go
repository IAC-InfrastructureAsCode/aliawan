package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Config is defined all the configuration needed by the app
type Config struct {
	Region    string `env:"ALI_REGION" envDefault:"ap-southeast-1"`
	AccesKey  string `env:"ALI_ACCESS_KEY"`
	SecretKey string `env:"ALI_SECRET_KEY"`
}

// LoadConfig for get all the configuration from Env Variables
func LoadConfig() *Config {
	//Init Config from ENV Variable
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
