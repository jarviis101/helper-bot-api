package main

import (
	"helper_openai_bot/internal/app"
	"helper_openai_bot/pkg"
	"log"
)

func main() {
	config, err := pkg.ResolveConfig()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	bot, err := pkg.ResolveBot(config.Telegram)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	client := pkg.ResolveClient(config.OpenAI)

	application := app.CreateApplication(bot, client)
	application.Run()
}
