package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command/service"
	"helper_openai_bot/internal/model"
)

type CommandHandler interface {
	Handle(update tgbotapi.Update) (*model.CommandResponse, error)
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

func (h *handler) Handle(update tgbotapi.Update) (*model.CommandResponse, error) {
	clientCommand, err := h.commandResolver.ResolveByCommand(update.Message.Text)
	if err != nil {
		return nil, err
	}
	strategy := h.strategyResolver.Resolve(clientCommand)

	msg := strategy.Handle(update)

	return &model.CommandResponse{Msg: msg}, nil
}
