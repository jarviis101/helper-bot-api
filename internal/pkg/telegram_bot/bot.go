package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func ResolveBot() (*tgbotapi.BotAPI, error) {
	if bot != nil {
		bot.Debug = true
		return bot, nil
	}

	bot, err := tgbotapi.NewBotAPI("6161208907:AAFPg5vchq52ySjfn0etSvjWyU-Ji4ZjT6w")
	bot.Debug = true
	if err != nil {
		return nil, err
	}

	return bot, nil
}
