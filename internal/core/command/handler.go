package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command/service"
	"log"
)

type CommandHandler interface {
	Handle(update tgbotapi.Update)
}

type handler struct {
	commandResolver  service.CommandResolver
	strategyResolver service.CommandStrategyResolver
}

func CreateCommandHandler(
	commandResolver service.CommandResolver,
	strategyResolver service.CommandStrategyResolver,
) CommandHandler {
	return &handler{
		commandResolver,
		strategyResolver,
	}
}

func (h *handler) Handle(update tgbotapi.Update) {
	clientCommand, err := h.commandResolver.ResolveByCommand(update.Message.Text)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	strategy := h.strategyResolver.Resolve(clientCommand)

	strategy.Handle(update)
}
