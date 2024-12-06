package config

import (
	"os"

	"github.com/agunfir98/gobroker/lib"
	"github.com/joho/godotenv"
)

type Config struct {
	ApiPort     string
	RabbitMQURL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	lib.FailOnError(err, "failed to load env")

	rabbitURL := os.Getenv("RABBIT_URL")

	config := &Config{
		ApiPort:     ":8080",
		RabbitMQURL: rabbitURL,
	}

	return config
}
