package common

import (
	"github.com/skeletongo/word/model"
)

func ToWord(w *TranslatedWord) *model.Word {
	ret := &model.Word{
		Text:   w.Text,
		UK:     w.UK,
		USA:    w.USA,
		Weight: 0,
		State:  model.Normal,
	}
	for _, v := range w.Means {
		ret.Means = append(ret.Means, &model.Mean{
			Typ:  int(v.Typ),
			Text: v.Text,
		})
	}
	return ret
}

func ToTranslatedWord(w *model.Word) *TranslatedWord {
	ret := &TranslatedWord{
		Word: &Word{
			Text: w.Text,
		},
		UK:  w.UK,
		USA: w.USA,
	}
	for _, v := range w.Means {
		ret.Means = append(ret.Means, &Mean{
			Typ:  WordType(v.Typ),
			Text: v.Text,
		})
	}
	return ret
}
