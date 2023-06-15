package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CommandResponse struct {
	Msg tgbotapi.MessageConfig
}

type TextResponse struct {
	Msg tgbotapi.MessageConfig
}
