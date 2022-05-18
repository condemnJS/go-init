package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}
