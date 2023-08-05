package text

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"helper_openai_bot/internal/service/infrastructure"
	"log"
)

type sender struct {
	client *openai.Client
}

func CreateSender(c *openai.Client) infrastructure.SenderInterface {
	return &sender{c}
}

func (s *sender) Send(message string, responseMessage chan string) {
	resp, err := s.client.CreateChatCompletion(
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
