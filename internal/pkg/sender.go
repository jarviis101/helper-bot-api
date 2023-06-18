package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
)

type SenderInterface interface {
	SendMessage(update tgbotapi.Update, msg tgbotapi.MessageConfig)
}

type sender struct {
	bot         *tgbotapi.BotAPI
	localizer   *i18n.Localizer
	maintenance bool
}

func CreateSender(bot *tgbotapi.BotAPI, localizer *i18n.Localizer, maintenanceMode bool) SenderInterface {
	return &sender{bot, localizer, maintenanceMode}
}

func (s *sender) SendMessage(update tgbotapi.Update, msg tgbotapi.MessageConfig) {
	if s.maintenance {
		maintenanceMessage, _ := s.localizer.Localize(&i18n.LocalizeConfig{MessageID: "maintenance_mode_message"})
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, maintenanceMessage)
		s.sendMessage(msg)

		return
	}
	if msg.Text == "" {
		somethingWenWrongMessage, _ := s.localizer.Localize(
			&i18n.LocalizeConfig{MessageID: "something_went_wrong_message"},
		)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, somethingWenWrongMessage)
		s.sendMessage(msg)

		return
	}

	s.sendMessage(msg)
}

func (s *sender) sendMessage(msg tgbotapi.MessageConfig) {
	if _, err := s.bot.Send(msg); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
