package lyricusecase

import (
	"context"
	"math/rand"
	"os"
	"slices"

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

func (l *lyricUsecase) FetchByKey(ctx context.Context) ([]*graphql.Lyric, error) {
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

	var lyrics []*graphql.Lyric
	copyLyrics := make([]string, len(constants.EenenLyrics))
	copy(copyLyrics, constants.EenenLyrics)
	for i := 0; i < constants.LyricNumber-constants.FakeLyricNumber; i++ {
		// 歌詞をランダムに取り出す
		randNum := rand.Intn(len(constants.EenenLyrics))
		lyrics = append(lyrics, &graphql.Lyric{
			ID:    int64(i + 1),
			Lyric: copyLyrics[randNum],
		})
		// 一度使ったものは使わない
		copyLyrics = slices.Delete(copyLyrics, randNum, randNum+1)
	}
	lyrics = append(lyrics, &graphql.Lyric{
		ID:    int64(constants.LyricNumber),
		Lyric: result,
	})
	// このままだと4番目が必ず偽物なので混ぜる
	rand.Shuffle(len(lyrics), func(i, j int) {
		lyrics[i], lyrics[j] = lyrics[j], lyrics[i]
	})
	return lyrics, nil
}
