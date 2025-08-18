package service

import "url-shortener/internal/domain/model"

// Repository — интерфейс для слоя данных
type Repository interface {
	Save(url model.URL) error
	Find(id string) (model.URL, bool)
}

// URLService — бизнес-логика
type URLService struct {
	repo Repository
}

func NewURLService(repo Repository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) Shorten(original string, generateID func() string) (string, error) {
	id := generateID()
	url := model.URL{ID: id, OriginalURL: original}
	if err := s.repo.Save(url); err != nil {
		return "", err
	}
	return id, nil
}

func (s *URLService) Resolve(id string) (string, bool) {
	url, ok := s.repo.Find(id)
	if !ok {
		return "", false
	}
	return url.OriginalURL, true
}
