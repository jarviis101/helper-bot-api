package infrastructure

import "helper_openai_bot/internal/service/infrastructure"

type (
	Application interface {
		Run()
	}
	Container interface {
		ProvideHandler() infrastructure.HandlerService
	}
)
