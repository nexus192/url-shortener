package service

import (
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	"url-shortener/pkg/shortener"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Shorten(original string) (string, error) {
	id := shortener.GenerateID()
	url := model.URL{ID: id, Original: original}
	err := s.repo.Save(url)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) Resolve(id string) (string, bool) {
	url, ok := s.repo.Find(id)
	if !ok {
		return "", false
	}
	return url.Original, true
}
