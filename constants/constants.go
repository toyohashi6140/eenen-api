package constants

import (
	"fmt"
	"strings"
)

const (
	FakeLyricNumber int = 1
	LyricNumber     int = 4
)

var (
	EenenLyrics []string = []string{
		"何も言わんでもええねん",
		"何もせんでもええねん",
		"笑い飛ばせばええねん",
		"好きにするのがええねん",
		"感じるだけでええねん",
		"気持ち良ければええねん",
		"後悔してもええねん",
		"また始めたらええねん",
		"失敗してもええねん",
		"もう一回やったらええねん",
		"前を向いたらええねん",
		"胸を張ったらええねん",
		"僕はお前がええねん",
		"好きでいれたらええねん",
		"同じ夢を見れたらええねん",
		"そんなステキな二人がええねん",
		"心配せんでええねん",
		"僕を見てればええねん",
		"アイデアなんかええねん",
		"別になくてもええねん",
		"ハッタリだけでええねん",
		"背伸びしたってええねん",
		"カッときたってええねん",
		"終わりよければええねん",
		"そんな自分が好きならええねん",
		"そんな日々が好きならええねん",
		"情けなくてもええねん",
		"叫んでみればええねん",
		"苦い涙もええねん",
		"ポロリこぼれてええねん",
		"ちょっと休めばええねん",
		"フッと笑えばええねん",
		"何もなくてもええねん",
		"信じていればええねん",
		"意味がなくてもええねん",
		"何かを感じていればええねん",
		"他に何もいらんねん",
		"それでええねん",
	}
	Prompt = fmt.Sprintf(`
		「ええねん」という楽曲があります。
		この曲は、〇〇ええねんという言い回しが多数存在します。
		以下はその例です。
		%s
		このような⚪︎⚪︎ええねんという上記に類似する言葉を生成してください。
		ただし、応答は「⚪︎⚪︎ええねん」という単語のみで、個数は%dつでお願いします。
		また、例と同じものを生成してはいけません。
	`, strings.Join(EenenLyrics, "\n"), FakeLyricNumber)
)
