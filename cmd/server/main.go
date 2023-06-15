package main

import "helper_openai_bot/internal/app"

func main() {
	application := app.CreateApplication()
	application.Run()
}
