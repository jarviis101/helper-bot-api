package service

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/pkg"
	"helper_openai_bot/internal/service/command/infrastructure"
	interfaces "helper_openai_bot/internal/service/infrastructure"
)

type commandHandler struct {
	commandResolver  infrastructure.CommandResolver
	strategyResolver infrastructure.CommandStrategyResolver
	sender           pkg.SenderInterface
}

func CreateCommandHandler(
	cr infrastructure.CommandResolver,
	sr infrastructure.CommandStrategyResolver,
	s pkg.SenderInterface,
) interfaces.HandlerService {
	return &commandHandler{cr, sr, s}
}

func (h *commandHandler) Handle(update tgbotapi.Update) {
	clientCommand, err := h.commandResolver.ResolveByCommand(update.Message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n", err.Error()))
		h.sender.SendMessage(update, msg)

		return
	}

	strategy := h.strategyResolver.Resolve(clientCommand)
	strategy.Handle(update)
}
