package text

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/pkg"
	"log"
)

type TextHandler interface {
	Handle(update tgbotapi.Update)
}

type textHandler struct {
	client *openai.Client
	sender pkg.SenderInterface
}

func CreateTextHandler(client *openai.Client, sender pkg.SenderInterface) TextHandler {
	return &textHandler{client, sender}
}

func (t *textHandler) Handle(update tgbotapi.Update) {
	responseMessage := make(chan string)
	go t.handleText(update.Message.Text, responseMessage)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, <-responseMessage)
	t.sender.SendMessage(update, msg)
}

func (t *textHandler) handleText(message string, responseMessage chan string) {
	resp, err := t.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		responseMessage <- ""
		return
	}

	responseMessage <- resp.Choices[0].Message.Content
}
