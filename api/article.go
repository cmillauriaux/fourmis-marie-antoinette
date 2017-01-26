package ants

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Article struct {
	DateTime        int64
	ID              string
	Title           string
	Author          string
	Content         string `datastore:",noindex"`
	Published       bool
	PictureFileName string
}

func (a *Article) transformContentToHTML() string {
	unsafe := blackfriday.MarkdownCommon([]byte(a.Content))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(html)
}
