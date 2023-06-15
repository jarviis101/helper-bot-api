package strategy

import "helper_openai_bot/internal/core/command/infrastructure"

type donateHandler struct {
}

func CreateDonateHandler() CommandHandler {
	return &donateHandler{}
}

func (d donateHandler) Support(supportCommand *infrastructure.Command) bool {
	return *supportCommand == infrastructure.DonateCommand
}

func (d donateHandler) Handle() (string, error) {
	return "Привет, можешь задонатить мне немного денег, для будущего развития бота - 5168752080882056 :)", nil
}
