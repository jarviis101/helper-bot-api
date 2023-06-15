package strategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command/infrastructure"
)

type donateHandler struct {
}

func CreateDonateHandler() CommandHandler {
	return &donateHandler{}
}

func (d donateHandler) Support(supportCommand *infrastructure.Command) bool {
	return *supportCommand == infrastructure.DonateCommand
}

func (d donateHandler) Handle(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Hi, You can support this bot - 5168752080882056 :)",
	)
}
