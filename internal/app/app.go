package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/container"
	interfaces "helper_openai_bot/internal/infrastructure"
)

const (
	timeout = 60
)

type application struct {
	bot         *tgbotapi.BotAPI
	client      *openai.Client
	Maintenance bool
}

func CreateApplication(bot *tgbotapi.BotAPI, client *openai.Client, maintenance bool) interfaces.Application {
	return &application{bot, client, maintenance}
}

func (app *application) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	di := container.CreateContainer(app.bot, app.client, app.Maintenance)
	handler := di.ProvideHandler()

	for update := range app.bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}

		handler.Handle(update)
	}
}
