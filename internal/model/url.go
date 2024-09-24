package model

import "time"

type URL struct {
	ID          int
	ShortURL    string
	OriginalURL string
	CreatedAt   time.Time
}
