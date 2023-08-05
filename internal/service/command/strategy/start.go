package strategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"helper_openai_bot/internal/pkg"
	"helper_openai_bot/internal/service/command/infrastructure"
)

type startHandler struct {
	sender    pkg.SenderInterface
	localizer *i18n.Localizer
}

func CreateStartHandler(sender pkg.SenderInterface, localizer *i18n.Localizer) infrastructure.CommandHandler {
	return &startHandler{sender, localizer}
}

func (s *startHandler) Support(supportCommand *infrastructure.Command) bool {
	return *supportCommand == infrastructure.StartCommand
}

func (s *startHandler) Handle(update tgbotapi.Update) {
	startMessage, _ := s.localizer.Localize(&i18n.LocalizeConfig{MessageID: "start_message"})
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMessage)

	s.sender.SendMessage(update, msg)
}
