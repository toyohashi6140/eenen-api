package openai

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type (
	OpenAIConnector interface {
		Chat(context.Context) (openai.ChatCompletionResponse, error)
	}
	openAIClient struct {
		client   *openai.Client
		model    string
		messages []openai.ChatCompletionMessage
	}
)

func NewOpenAIConnector(token string, model string, messsages []openai.ChatCompletionMessage) OpenAIConnector {
	return &openAIClient{
		client:   openai.NewClient(token),
		model:    model,
		messages: messsages,
	}
}

func (o *openAIClient) Chat(ctx context.Context) (openai.ChatCompletionResponse, error) {
	return o.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    o.model,
		Messages: o.messages,
	})

}
