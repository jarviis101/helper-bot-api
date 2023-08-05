package command

import (
	"helper_openai_bot/internal/service/command/infrastructure"
)

type commandStrategyResolver struct {
	handleStrategies []infrastructure.CommandHandler
}

func CreateCommandStrategyResolver(handlers []infrastructure.CommandHandler) infrastructure.CommandStrategyResolver {
	return &commandStrategyResolver{handleStrategies: handlers}
}

func (s *commandStrategyResolver) Resolve(command *infrastructure.Command) infrastructure.CommandHandler {
	for _, handler := range s.handleStrategies {
		if handler.Support(command) {
			return handler
		}
	}

	return nil
}
