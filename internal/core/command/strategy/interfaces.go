package strategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command/infrastructure"
)

type (
	CommandHandler interface {
		Support(supportCommand *infrastructure.Command) bool
		Handle(update tgbotapi.Update) tgbotapi.MessageConfig
	}
)
