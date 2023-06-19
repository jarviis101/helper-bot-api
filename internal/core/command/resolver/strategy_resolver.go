package resolver

import (
	"helper_openai_bot/internal/core/command/infrastructure"
	"helper_openai_bot/internal/core/command/strategy"
)

type CommandStrategyResolver interface {
	Resolve(command *infrastructure.Command) strategy.CommandHandler
}

type commandStrategyResolver struct {
	handleStrategies []strategy.CommandHandler
}

func CreateCommandStrategyResolver(handlers []strategy.CommandHandler) CommandStrategyResolver {
	return &commandStrategyResolver{handleStrategies: handlers}
}

func (s *commandStrategyResolver) Resolve(command *infrastructure.Command) strategy.CommandHandler {
	for _, handler := range s.handleStrategies {
		if handler.Support(command) {
			return handler
		}
	}

	return nil
}
