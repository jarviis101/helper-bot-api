package text

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/model"
)

type TextHandler interface {
	Handle(update tgbotapi.Update) (*model.TextResponse, error)
}

type textHandler struct {
	client *openai.Client
}

func CreateTextHandler(client *openai.Client) TextHandler {
	return &textHandler{client}
}

func (t *textHandler) Handle(update tgbotapi.Update) (*model.TextResponse, error) {
	responseMessage, err := t.handleText(update.Message.Text)
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMessage)
	return &model.TextResponse{Msg: msg}, nil
}

func (t *textHandler) handleText(message string) (string, error) {
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
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
