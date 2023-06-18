package pkg

import "github.com/sashabaranov/go-openai"

func CreateClient(ai OpenAI) *openai.Client {
	return openai.NewClient(ai.Token)
}
