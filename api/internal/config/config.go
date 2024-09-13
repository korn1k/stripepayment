package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Stripe stripe
}

type stripe struct {
	SecretKey string
}

func MustInitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load .env")
	}

	return &Config{
		Stripe: stripe{SecretKey: os.Getenv("STRIPE_SECRET_KEY")},
	}
}