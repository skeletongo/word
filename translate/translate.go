package translate

import (
	"github.com/skeletongo/word/common"
)

type Translate struct {
}

func (t *Translate) Translate(word *common.Word) *common.TranslatedWord {
	// todo
	return &common.TranslatedWord{
		Word: word,
		UK:   "həˈləʊ",
		USA:  "həˈləʊ",
		Means: []*common.Mean{
			{
				Typ:  common.Noun,
				Text: "招呼，问候；（Hello）（法、印、美、俄）埃洛（人名）",
			},
			{
				Typ:  common.Verb,
				Text: "说（或大声说）“喂”；打招呼",
			},
			{
				Typ:  common.Interjection,
				Text: "喂，你好（用于问候或打招呼）；喂，你好（打电话时的招呼语）；喂，你好（引起别人注意的招呼语）；<非正式>喂，嘿 (认为别人说了蠢话或分心)；<英，旧>嘿（表示惊讶）",
			},
		},
	}
}

func NewTranslate() *Translate {
	return &Translate{}
}
