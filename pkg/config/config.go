package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/kalougata/go-take-out/configs"
)

type Config struct {
	DB configs.Database
}

func NewConfig() *Config {

	var config string

	flag.StringVar(&config, "config", ".env", "config file, eg: --config=[.filename]")
	flag.Parse()

	if err := godotenv.Load(config); err != nil {
		log.Fatalf("failed to load env file: %s", err)
	}

	return &Config{
		DB: configs.DatabaseConfig(),
	}
}
