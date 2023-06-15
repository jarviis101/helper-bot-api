package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/core/command"
	"helper_openai_bot/internal/core/command/service"
	"helper_openai_bot/internal/core/text"
)

type Handler interface {
	Handle(update tgbotapi.Update) tgbotapi.MessageConfig
}

type handler struct {
	commandStrategyResolver service.CommandStrategyResolver
	commandResolver         service.CommandResolver
	openAIClient            *openai.Client
}

func CreateHandler(
	commandStrategyResolver service.CommandStrategyResolver,
	commandResolver service.CommandResolver,
	openAIClient *openai.Client,
) Handler {
	return &handler{
		commandStrategyResolver,
		commandResolver,
		openAIClient,
	}
}

func (h *handler) Handle(update tgbotapi.Update) tgbotapi.MessageConfig {
	if update.Message.IsCommand() {
		return h.handleCommand(update)
	}

	return h.handleTextMessage(update)
}

func (h *handler) handleCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	commandHandler := command.CreateCommandHandler(h.commandResolver, h.commandStrategyResolver)
	commandModel, err := commandHandler.Handle(update)
	msg := commandModel.Msg
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	return msg
}

func (h *handler) handleTextMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	textHandler := text.CreateTextHandler(h.openAIClient)
	responseModel, err := textHandler.Handle(update)
	msg := responseModel.Msg
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	return msg
}
