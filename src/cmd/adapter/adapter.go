package adapter

import (
	"url-shortener/internal/data"
	"url-shortener/internal/data/repository"
	"url-shortener/internal/domain/model"
)

type PostgresAdapter struct {
	Repo *repository.PostgresRepo
}

func (a *PostgresAdapter) Save(url model.URL) error {
	return a.Repo.Save(data.ToData(url))
}

func (a *PostgresAdapter) Find(id string) (model.URL, bool) {
	u, ok := a.Repo.Find(id)
	if !ok {
		return model.URL{}, false
	}
	return data.ToDomain(u), true
}
