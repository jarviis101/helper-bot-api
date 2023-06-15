package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func ResolveBot(telegram Telegram) (*tgbotapi.BotAPI, error) {
	if bot != nil {
		bot.Debug = true
		return bot, nil
	}

	telegramBot, err := tgbotapi.NewBotAPI(telegram.Token)
	telegramBot.Debug = true
	if err != nil {
		return nil, err
	}

	return telegramBot, nil
}
