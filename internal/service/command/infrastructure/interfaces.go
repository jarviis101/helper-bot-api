package infrastructure

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	CommandHandler interface {
		Support(supportCommand *Command) bool
		Handle(update tgbotapi.Update)
	}
	CommandResolver interface {
		ResolveByCommand(rawCommand string) (*Command, error)
	}
	CommandStrategyResolver interface {
		Resolve(command *Command) CommandHandler
	}
)
