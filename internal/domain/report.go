package domain

import "time"

type Report struct {
	ID         int
	IDPosts    int `json:",omitempty"`
	IDComments int `json:",omitempty"`
	UserName   string
	Text       string
	CreateDate time.Time
}
