package repository

import "url-shortener/internal/domain/model"

type InMemoryRepo struct {
	data map[string]model.URL
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{data: make(map[string]model.URL)}
}

func (r *InMemoryRepo) Save(url model.URL) error {
	r.data[url.ID] = url
	return nil
}

func (r *InMemoryRepo) Find(id string) (model.URL, bool) {
	url, ok := r.data[id]
	return url, ok
}
