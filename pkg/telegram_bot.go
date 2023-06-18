package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateBot(telegram Telegram, maintenance bool) (*tgbotapi.BotAPI, error) {
	telegramBot, err := tgbotapi.NewBotAPI(telegram.Token)
	telegramBot.Debug = maintenance
	if err != nil {
		return nil, err
	}

	return telegramBot, nil
}
