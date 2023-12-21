package openaiutils

import openai "github.com/sashabaranov/go-openai"

func GetFirstChatMessage(res openai.ChatCompletionResponse) string {
	if len(res.Choices) == 0 {
		return ""
	}
	return res.Choices[0].Message.Content
}
