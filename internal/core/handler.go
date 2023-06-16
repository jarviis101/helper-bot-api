package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command"
	"helper_openai_bot/internal/core/text"
)

type Handler interface {
	Handle(update tgbotapi.Update) tgbotapi.MessageConfig
}

type handler struct {
	commandHandler command.CommandHandler
	textHandler    text.TextHandler
}

func CreateHandler(commandHandler command.CommandHandler, textHandler text.TextHandler) Handler {
	return &handler{
		commandHandler,
		textHandler,
	}
}

func (h *handler) Handle(update tgbotapi.Update) tgbotapi.MessageConfig {
	if update.Message.IsCommand() {
		return h.handleCommand(update)
	}

	return h.handleTextMessage(update)
}

func (h *handler) handleCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	commandModel, err := h.commandHandler.Handle(update)
	if err != nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	}

	return commandModel.Msg
}

func (h *handler) handleTextMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	responseModel := h.textHandler.Handle(update)

	return responseModel.Msg
}
