package infrastructure

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type (
	HandlerService interface {
		Handle(update tgbotapi.Update)
	}
	SenderInterface interface {
		Send(message string, responseMessage chan string)
	}
)
