package container

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/controller"
	lowinterfaces "helper_openai_bot/internal/infrastructure"
	"helper_openai_bot/internal/pkg"
	"helper_openai_bot/internal/service"
	"helper_openai_bot/internal/service/command"
	highinterfaces "helper_openai_bot/internal/service/command/infrastructure"
	"helper_openai_bot/internal/service/command/strategy"
	"helper_openai_bot/internal/service/infrastructure"
	"helper_openai_bot/internal/service/text"
)

type container struct {
	bot         *tgbotapi.BotAPI
	client      *openai.Client
	Maintenance bool
}

func CreateContainer(bot *tgbotapi.BotAPI, client *openai.Client, maintenance bool) lowinterfaces.Container {
	return &container{bot, client, maintenance}
}

func (c *container) ProvideHandler() infrastructure.HandlerService {
	localizer := pkg.CreateLocalizer()
	messageSender := pkg.CreateSender(c.bot, localizer, c.Maintenance)

	commandStrategyResolver := c.resolveCommandStrategyResolver(localizer, messageSender)
	commandResolver := command.CreateCommandResolver()
	openaiSender := text.CreateSender(c.client)

	commandHandler := service.CreateCommandHandler(commandResolver, commandStrategyResolver, messageSender)
	textHandler := service.CreateTextHandler(messageSender, openaiSender)

	return controller.CreateHandler(commandHandler, textHandler)
}

func (c *container) resolveCommandStrategyResolver(
	localizer *i18n.Localizer,
	sender pkg.SenderInterface,
) highinterfaces.CommandStrategyResolver {
	startHandler := strategy.CreateStartHandler(sender, localizer)
	donateHandler := strategy.CreateDonateHandler(sender, localizer)

	return command.CreateCommandStrategyResolver([]highinterfaces.CommandHandler{startHandler, donateHandler})
}
