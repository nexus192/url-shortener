package model

import (
	"time"
)

type URLData struct {
	ID        string
	URL       string
	Short     string
	CreatedAt time.Time
}
