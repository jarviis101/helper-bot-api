package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/service/infrastructure"
)

type handler struct {
	commandHandler infrastructure.HandlerService
	textHandler    infrastructure.HandlerService
}

func CreateHandler(ch, th infrastructure.HandlerService) infrastructure.HandlerService {
	return &handler{ch, th}
}

func (h *handler) Handle(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		h.commandHandler.Handle(update)
		return
	}

	h.textHandler.Handle(update)
}
