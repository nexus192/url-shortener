package main

import (
	"log"
	"net/http"

	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	repo := repository.NewInMemoryRepo()
	svc := service.New(repo)
	h := handler.New(svc)

	r := chi.NewRouter()
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{id}", h.Redirect)

	log.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
