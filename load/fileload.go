package load

import "github.com/skeletongo/word/common"

type FileLoad struct {
}

func (u *FileLoad) FromFile(url string) []*common.Word {
	//TODO implement me
	panic("implement me")
}

func NewFileLoad() *FileLoad {
	return &FileLoad{}
}
