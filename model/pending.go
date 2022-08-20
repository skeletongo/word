package model

import "gorm.io/gorm"

// Pending 待翻译英文单词表
type Pending struct {
	gorm.Model
	Text  string `gorm:"uniqueIndex"` // 英文单词
	State int    // 状态
}
