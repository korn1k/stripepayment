package config

import ("os")

type Config struct {
	Stripe stripe
}

type stripe struct {
	SecretKey string
}

func MustInitConfig() *Config {
	return &Config{
		Stripe: stripe{SecretKey: os.Getenv("STRIPE_SECRET_KEY")},
	}
}