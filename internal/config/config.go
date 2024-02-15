package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	var cfg Config

	cfg.DBUrl = os.Getenv("DB_URL")
	cfg.Port = os.Getenv("PORT")

	return &cfg, nil
}
