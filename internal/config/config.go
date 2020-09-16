package config

// if using go modules

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port             int    `env:"Port" envDefault:"1323"`
	GrpcPort         string `env:"GrpcPort"`
	DbDriver         string `env:"DbDriver"`
	ConStr           string `env:"ConStr"`
}

func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &cfg
}
