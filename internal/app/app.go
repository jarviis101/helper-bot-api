package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core"
	"helper_openai_bot/internal/core/command/service"
	"helper_openai_bot/internal/core/command/strategy"
	"helper_openai_bot/internal/pkg"
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
}

func CreateApplication() Application {
	return &application{}
}

func (app *application) Run() {
	config, err := pkg.ResolveConfig()
	if err != nil {
		log.Println(err.Error())
	}
	bot, err := pkg.ResolveBot(config.Telegram)
	if err != nil {
		log.Println(err.Error())
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	openAIClient := pkg.ResolveClient(config.OpenAI)

	startHandler := strategy.CreateStartHandler()
	donateHandler := strategy.CreateDonateHandler()
	commandStrategyResolver := service.CreateCommandStrategyResolver([]strategy.CommandHandler{startHandler, donateHandler})
	commandResolver := service.CreateCommandResolver()

	handler := core.CreateHandler(commandStrategyResolver, commandResolver, openAIClient)

	for update := range bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}

		msg := handler.Handle(update)
		app.sendMessage(bot, msg)
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
