package common

// 业务逻辑相关常量

type WordType int

func (w WordType) String() string {
	return WordTypeToStringMap[w]
}

func (w WordType) Short() string {
	return WordTypeToShortMap[w]
}

const (
	Noun             WordType = iota + 1 // 名词
	Pronoun                              // 代词
	Adjective                            // 形容词
	Adverb                               // 副词
	Verb                                 // 动词
	VerbIntransitive                     // 不及物动词
	VerbTransitive                       // 及物动词
	Auxiliary                            // 助动词
	Numeral                              // 数词
	Article                              // 冠词
	Preposition                          // 介词
	Conjunction                          // 连词
	Interjection                         // 感叹词

)

var WordTypeToStringMap = map[WordType]string{
	Noun:             "名词",
	Pronoun:          "代词",
	Adjective:        "形容词",
	Adverb:           "副词",
	Verb:             "动词",
	VerbIntransitive: "不及物动词",
	VerbTransitive:   "及物动词",
	Auxiliary:        "助动词",
	Numeral:          "数词",
	Article:          "冠词",
	Preposition:      "介词",
	Conjunction:      "连词",
	Interjection:     "感叹词",
}

var WordTypeToShortMap = map[WordType]string{
	Noun:             "n",
	Pronoun:          "pron",
	Adjective:        "adj",
	Adverb:           "adv",
	Verb:             "v",
	VerbIntransitive: "vi",
	VerbTransitive:   "vt",
	Auxiliary:        "aux",
	Numeral:          "num",
	Article:          "art",
	Preposition:      "prep",
	Conjunction:      "conj",
	Interjection:     "int",
}

var StringToWordTypeMap = map[string]WordType{}
var ShortToWordTypeMap = map[string]WordType{}

func init() {
	for k, v := range WordTypeToStringMap {
		StringToWordTypeMap[v] = k
	}
	for k, v := range WordTypeToShortMap {
		ShortToWordTypeMap[v] = k
	}
}
