package lyricusecase

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/toyohashi6140/eenen-api/constants"
	"github.com/toyohashi6140/eenen-api/domain/models/graphql"
	myopenai "github.com/toyohashi6140/eenen-api/openai"
	"github.com/toyohashi6140/eenen-api/openai/openaiutils"
	"github.com/toyohashi6140/eenen-api/usecases"
)

type (
	LyricUsecase interface {
		usecases.ManyReader[*graphql.Lyric]
	}
	lyricUsecase struct {
		// converter
	}
)

func NewLyricUsecase() LyricUsecase {
	return &lyricUsecase{}
}

func (l *lyricUsecase) FetchByKey(ctx context.Context) (*graphql.Lyric, error) {
	ai := myopenai.NewOpenAIConnector(os.Getenv("OPEN_AI_API_KEY"), openai.GPT3Dot5Turbo, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: constants.Prompt,
		},
	})
	res, err := ai.Chat(ctx)
	if err != nil {
		return nil, err
	}
	result := openaiutils.GetFirstChatMessage(res)
	return &graphql.Lyric{
		ID:    1,
		Lyric: result,
	}, nil
}
