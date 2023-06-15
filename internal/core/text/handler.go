package text

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/model"
	openai_client "helper_openai_bot/internal/pkg/openai"
)

type TextHandler interface {
	Handle(c string) (*model.TextResponse, error)
}

type textHandler struct {
	client *openai.Client
}

func CreateTextHandler() TextHandler {
	client := openai_client.ResolveClient()
	return &textHandler{client}
}

func (t *textHandler) Handle(message string) (*model.TextResponse, error) {
	responseMessage, err := t.handleText(message)
	if err != nil {
		return nil, err
	}

	return &model.TextResponse{Message: responseMessage}, nil
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
