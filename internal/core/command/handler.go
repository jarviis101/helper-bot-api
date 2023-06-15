package command

import (
	"helper_openai_bot/internal/core/command/service"
	"helper_openai_bot/internal/model"
)

type CommandHandler interface {
	Handle(c string) (*model.CommandResponse, error)
}

type handler struct {
	commandResolver  service.CommandResolver
	strategyResolver service.CommandStrategyResolver
}

func CreateCommandHandler(commandResolver service.CommandResolver, strategyResolver service.CommandStrategyResolver) CommandHandler {
	return &handler{
		commandResolver,
		strategyResolver,
	}
}

func (h *handler) Handle(c string) (*model.CommandResponse, error) {
	clientCommand, err := h.commandResolver.ResolveByCommand(c)
	if err != nil {
		return nil, err
	}
	strategy := h.strategyResolver.Resolve(clientCommand)

	message, err := strategy.Handle()
	if err != nil {
		return nil, err
	}

	return &model.CommandResponse{Message: message}, nil
}
