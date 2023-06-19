package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/core"
	"helper_openai_bot/internal/core/command"
	"helper_openai_bot/internal/core/command/resolver"
	"helper_openai_bot/internal/core/command/strategy"
	"helper_openai_bot/internal/core/text"
	"helper_openai_bot/internal/pkg"
)

const (
	timeout = 60
)

type Application interface {
	Run()
}

type application struct {
	bot         *tgbotapi.BotAPI
	client      *openai.Client
	Maintenance bool
}

func CreateApplication(bot *tgbotapi.BotAPI, client *openai.Client, maintenance bool) Application {
	return &application{bot, client, maintenance}
}

func (app *application) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	localizer := pkg.CreateLocalizer()
	sender := pkg.CreateSender(app.bot, localizer, app.Maintenance)

	commandStrategyResolver := app.resolveCommandStrategyResolver(localizer, sender)
	commandResolver := resolver.CreateCommandResolver()

	commandHandler := command.CreateCommandHandler(commandResolver, commandStrategyResolver)
	textHandler := text.CreateTextHandler(app.client, sender)

	handler := core.CreateHandler(commandHandler, textHandler)

	for update := range app.bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}

		handler.Handle(update)
	}
}

func (app *application) resolveCommandStrategyResolver(
	localizer *i18n.Localizer,
	sender pkg.SenderInterface,
) resolver.CommandStrategyResolver {
	startHandler := strategy.CreateStartHandler(sender, localizer)
	donateHandler := strategy.CreateDonateHandler(sender, localizer)

	return resolver.CreateCommandStrategyResolver([]strategy.CommandHandler{startHandler, donateHandler})
}
