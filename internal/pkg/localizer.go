package pkg

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

func CreateLocalizer() *i18n.Localizer {
	bundle := i18n.NewBundle(language.Ukrainian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	if _, err := bundle.LoadMessageFile("resources/uk.json"); err != nil {
		log.Printf("Error: %s\n", err.Error())

		return nil
	}

	return i18n.NewLocalizer(bundle, language.Ukrainian.String())
}
