package postmodel

import "time"

type PostResponse struct {
	Id          string     `json:"id" `
	Title       string     `json:"title" `
	Content     string     `json:"content" `
	Published   bool       `json:"published" `
	PublishedAt *time.Time `json:"published_at"`
}
