package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    TelegramToken    string
    OpenRouterKey    string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    return &Config{
        TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
        OpenRouterKey: os.Getenv("OPENROUTER_API_KEY"),
    }, nil
} 