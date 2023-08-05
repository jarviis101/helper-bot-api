package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/pkg"
	"helper_openai_bot/internal/service/infrastructure"
)

type textHandler struct {
	messageSender pkg.SenderInterface
	openaiSender  infrastructure.SenderInterface
}

func CreateTextHandler(s pkg.SenderInterface, h infrastructure.SenderInterface) infrastructure.HandlerService {
	return &textHandler{s, h}
}

func (t *textHandler) Handle(update tgbotapi.Update) {
	responseMessage := make(chan string)
	go t.openaiSender.Send(update.Message.Text, responseMessage)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, <-responseMessage)
	t.messageSender.SendMessage(update, msg)
}
