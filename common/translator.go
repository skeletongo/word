package common

type Word struct {
	Text string // 英文单词
}

type Mean struct {
	Typ  WordType // 词性
	Text string   // 译文
}

type TranslatedWord struct {
	*Word  // 英文单词
	UK     string
	USA    string
	Means  []*Mean // 翻译结果
	Weight int     // 权重
	State  int     // 状态
}

type Translator interface {
	Translate(word *Word) *TranslatedWord
}
