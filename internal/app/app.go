package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/core"
	"helper_openai_bot/internal/core/command"
	"helper_openai_bot/internal/core/command/service"
	"helper_openai_bot/internal/core/command/strategy"
	"helper_openai_bot/internal/core/text"
	"log"
)

const (
	maintenance = false
	timeout     = 60
)

type Application interface {
	Run()
}

type application struct {
	bot    *tgbotapi.BotAPI
	client *openai.Client
}

func CreateApplication(bot *tgbotapi.BotAPI, client *openai.Client) Application {
	return &application{bot, client}
}

func (app *application) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	commandStrategyResolver := app.resolveCommandStrategyResolver()
	commandResolver := service.CreateCommandResolver()

	commandHandler := command.CreateCommandHandler(commandResolver, commandStrategyResolver)
	textHandler := text.CreateTextHandler(app.client)

	handler := core.CreateHandler(commandHandler, textHandler)

	for update := range app.bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}

		msg := handler.Handle(update)
		app.sendMessage(app.bot, msg)
	}
}

func (app *application) sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if maintenance {
		msg.Text = "Bot in maintenance mode"
	}
	if msg.Text == "" {
		msg.Text = "Что то пошло не так, ознакомься со списком возможностей бота :)"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Println(err.Error())
	}
}

func (app *application) resolveCommandStrategyResolver() service.CommandStrategyResolver {
	startHandler := strategy.CreateStartHandler()
	donateHandler := strategy.CreateDonateHandler()
	return service.CreateCommandStrategyResolver([]strategy.CommandHandler{startHandler, donateHandler})
}
