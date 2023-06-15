package strategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command/infrastructure"
)

type startHandler struct {
}

func CreateStartHandler() CommandHandler {
	return &startHandler{}
}

func (s *startHandler) Support(supportCommand *infrastructure.Command) bool {
	return *supportCommand == infrastructure.StartCommand
}

func (s *startHandler) Handle(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, можешь задать мне любой вопрос :)")
}
