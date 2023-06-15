package strategy

import (
	"helper_openai_bot/internal/core/command/infrastructure"
)

type startHandler struct {
}

func CreateStartHandler() CommandHandler {
	return &startHandler{}
}

func (s *startHandler) Support(supportCommand *infrastructure.Command) bool {
	return *supportCommand == infrastructure.StartCommand
}

func (s *startHandler) Handle() (string, error) {
	return "Привет, можешь задать мне любой вопрос :)", nil
}
