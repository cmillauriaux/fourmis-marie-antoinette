package ants

import (
	"time"
)

type Article struct {
	DateTime        time.Time
	ID              string
	Title           string
	Author          string
	Content         string
	Published       bool
	PictureFileName string
}
