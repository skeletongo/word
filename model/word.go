package model

import (
	"gorm.io/gorm"
)

const (
	Normal = iota
	Delete
)

type Mean struct {
	Typ  int    // 词性
	Text string // 译文
}

// Word 英文及中文翻译表
type Word struct {
	gorm.Model
	Text   string  `gorm:"uniqueIndex"` // 英文单词
	UK     string  // 英式音标
	USA    string  // 美式音标
	Means  []*Mean `gorm:"serializer:json"` // 中文翻译
	Weight int     // 权重
	State  int     // 状态
}
