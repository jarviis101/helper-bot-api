package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/core/command"
	"helper_openai_bot/internal/core/command/service"
	"helper_openai_bot/internal/core/text"
)

type Handler interface {
	Handle(update tgbotapi.Update) *tgbotapi.MessageConfig
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

func (h *handler) Handle(update tgbotapi.Update) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.IsCommand() {
		return h.handleCommand(&msg, update.Message.Text)
	}

	return h.handleTextMessage(&msg, update.Message.Text)
}

func (h *handler) handleCommand(msg *tgbotapi.MessageConfig, clientCommand string) *tgbotapi.MessageConfig {
	commandHandler := command.CreateCommandHandler(h.commandResolver, h.commandStrategyResolver)
	commandModel, err := commandHandler.Handle(clientCommand)
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	msg.Text = commandModel.Message
	return msg
}

func (h *handler) handleTextMessage(msg *tgbotapi.MessageConfig, clientMessage string) *tgbotapi.MessageConfig {
	textHandler := text.CreateTextHandler(h.openAIClient)
	responseModel, err := textHandler.Handle(clientMessage)
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	msg.Text = responseModel.Message
	return msg
}
