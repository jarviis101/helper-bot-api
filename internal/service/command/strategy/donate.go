package strategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"helper_openai_bot/internal/pkg"
	"helper_openai_bot/internal/service/command/infrastructure"
)

type donateHandler struct {
	sender    pkg.SenderInterface
	localizer *i18n.Localizer
}

func CreateDonateHandler(sender pkg.SenderInterface, localizer *i18n.Localizer) infrastructure.CommandHandler {
	return &donateHandler{sender, localizer}
}

func (d donateHandler) Support(supportCommand *infrastructure.Command) bool {
	return *supportCommand == infrastructure.DonateCommand
}

func (d donateHandler) Handle(update tgbotapi.Update) {
	startMessage, _ := d.localizer.Localize(&i18n.LocalizeConfig{MessageID: "support_message"})
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMessage)

	d.sender.SendMessage(update, msg)
}
