package strategy

import "helper_openai_bot/internal/core/command/infrastructure"

type (
	CommandHandler interface {
		Support(supportCommand *infrastructure.Command) bool
		Handle() (string, error)
	}
)
