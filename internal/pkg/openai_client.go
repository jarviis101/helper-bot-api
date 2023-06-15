package pkg

import "github.com/sashabaranov/go-openai"

var client *openai.Client

func ResolveClient(ai OpenAI) *openai.Client {
	if client != nil {
		return client
	}

	return openai.NewClient(ai.Token)
}
