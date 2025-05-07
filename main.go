package main

import (
    "log"

    "github.com/TatsianaHalaburda/Gen-AI-tg-bot/bot"
    "github.com/TatsianaHalaburda/Gen-AI-tg-bot/config"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading config:", err)
    }

    b, err := bot.NewBot(cfg.TelegramToken, cfg.OpenRouterKey)
    if err != nil {
        log.Fatal("Error creating bot:", err)
    }

    b.Start()
} 