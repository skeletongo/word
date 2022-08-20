package common

type FileLoader interface {
	FromFile(src string) []*Word
}

type UrlLoader interface {
	FromUrl(url string) []*Word
}
