package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	// Параметры для подключения к базе данных
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     int    `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
}

func MustLoad() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
