package model

import (
	"time"
	// domain "url-shortener/internal/domain/model"
)

type URLData struct {
	ID        string
	URL       string
	Short     string
	CreatedAt time.Time
}

// // Data -> Domain
// func (u URLData) ToDomain() domain.URL {
// 	return domain.URL{
// 		ID:       u.ID,
// 		Original: u.OriginalURL,
// 	}
// }

// // Domain -> Data
// func FromDomain(u domain.URL) URLData {
// 	return URLData{
// 		ID:          u.ID,
// 		OriginalURL: u.Original,
// 		CreatedAt:   time.Now(),
// 	}
// }
