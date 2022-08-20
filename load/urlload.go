package load

import (
	"github.com/skeletongo/word/common"
)

type UrlLoad struct {
}

func (u *UrlLoad) FromUrl(url string) []*common.Word {
	//TODO implement me
	panic("implement me")
}

func NewUrlLoad() *UrlLoad {
	return &UrlLoad{}
}
