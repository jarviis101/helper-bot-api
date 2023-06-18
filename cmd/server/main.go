package main

import (
	"helper_openai_bot/internal/app"
	"helper_openai_bot/pkg"
	"log"
)

const maintenance = false

func main() {
	config, err := pkg.CreateConfig()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	bot, err := pkg.CreateBot(config.Telegram, maintenance)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	client := pkg.CreateClient(config.OpenAI)

	application := app.CreateApplication(bot, client, maintenance)
	application.Run()
}
