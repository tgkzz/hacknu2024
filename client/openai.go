package client

import (
	"github.com/sashabaranov/go-openai"
)

func NewOpenaiClient(token string) *openai.Client {
	client := openai.NewClient(token)

	return client
}
