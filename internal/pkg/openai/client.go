package openai

import "github.com/sashabaranov/go-openai"

var client *openai.Client

func ResolveClient() *openai.Client {
	if client != nil {
		return client
	}

	return openai.NewClient("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")
}
