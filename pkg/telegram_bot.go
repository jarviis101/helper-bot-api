package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/app"
)

var bot *tgbotapi.BotAPI

func ResolveBot(telegram Telegram) (*tgbotapi.BotAPI, error) {
	if bot != nil {
		bot.Debug = app.Maintenance
		return bot, nil
	}

	telegramBot, err := tgbotapi.NewBotAPI(telegram.Token)
	telegramBot.Debug = app.Maintenance
	if err != nil {
		return nil, err
	}

	return telegramBot, nil
}
