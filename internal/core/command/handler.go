package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command/resolver"
	"log"
)

type Handler interface {
	Handle(update tgbotapi.Update)
}

type handler struct {
	commandResolver  resolver.CommandResolver
	strategyResolver resolver.CommandStrategyResolver
}

func CreateCommandHandler(
	commandResolver resolver.CommandResolver,
	strategyResolver resolver.CommandStrategyResolver,
) Handler {
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
