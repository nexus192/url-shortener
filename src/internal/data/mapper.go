package data

import (
	"time"
	dataModel "url-shortener/src/internal/data/model"
	domainModel "url-shortener/src/internal/domain/model"
)

func ToData(u domainModel.URL) dataModel.URLData {
	return dataModel.URLData{
		ID:        u.ID,
		URL:       u.OriginalURL,
		CreatedAt: time.Now(),
	}
}

func ToDomain(u dataModel.URLData) domainModel.URL {
	return domainModel.URL{
		ID:          u.ID,
		OriginalURL: u.URL,
	}
}
