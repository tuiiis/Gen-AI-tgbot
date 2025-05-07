package bot

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "gopkg.in/telebot.v3"
    "github.com/TatsianaHalaburda/Gen-AI-tg-bot/utils"
)

type Bot struct {
    bot            *telebot.Bot
    openRouterKey  string
    httpClient     *http.Client
}

type OpenRouterRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OpenRouterResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}

func NewBot(token string, openRouterKey string) (*Bot, error) {
    b, err := telebot.NewBot(telebot.Settings{
        Token: token,
    })
    if err != nil {
        return nil, err
    }

    return &Bot{
        bot:           b,
        openRouterKey: openRouterKey,
        httpClient:    &http.Client{},
    }, nil
}

func (b *Bot) Start() {
    b.bot.Handle("/start", func(c telebot.Context) error {
        return c.Send("Привет! Я бот, который может переворачивать сообщения и генерировать истории. Используйте /story для генерации истории.")
    })

    b.bot.Handle("/story", func(c telebot.Context) error {
        story, err := b.generateStory()
        if err != nil {
            return c.Send("Извините, произошла ошибка при генерации истории.")
        }
        return c.Send(story)
    })

    b.bot.Handle(telebot.OnText, func(c telebot.Context) error {
        reversed := utils.ReverseString(c.Text())
        return c.Send(reversed)
    })

    b.bot.Start()
}

func (b *Bot) generateStory() (string, error) {
    reqBody := OpenRouterRequest{
        Model: "meta-llama/llama-2-70b-chat",
        Messages: []Message{
            {
                Role:    "user",
                Content: "Напиши короткую научно-фантастическую историю не более 400 символов.",
            },
        },
    }

    jsonData, err := json.Marshal(reqBody)
    if err != nil {
        return "", fmt.Errorf("error marshaling request: %v", err)
    }

    req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Authorization", "Bearer "+b.openRouterKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("HTTP-Referer", "https://github.com/yourusername/telegram-bot")
    req.Header.Set("X-Title", "Telegram Bot")

    resp, err := b.httpClient.Do(req)
    if err != nil {
        return "", fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response: %v", err)
    }

    var openRouterResp OpenRouterResponse
    if err := json.Unmarshal(body, &openRouterResp); err != nil {
        return "", fmt.Errorf("error unmarshaling response: %v", err)
    }

    if len(openRouterResp.Choices) == 0 {
        return "", fmt.Errorf("no response from model")
    }

    return openRouterResp.Choices[0].Message.Content, nil
} 