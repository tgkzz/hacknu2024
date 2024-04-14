package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host        string
	Port        string
	MongoDB     MongoDB
	OpenAiToken string
}

type MongoDB struct {
	User     string
	Password string
	Host     string
	Name     string
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {

	}

	cfg := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		MongoDB: MongoDB{
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PSW"),
			Host:     os.Getenv("MONGO_HOST"),
			Name:     os.Getenv("MONGO_DB_NAME"),
		},
		OpenAiToken: os.Getenv("OPENAI_KEY"),
	}

	return cfg, nil
}
