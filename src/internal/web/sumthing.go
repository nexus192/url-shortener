package web

import (
	"html/template"
	"net/http"
	"url-shortener/internal/domain/service"
)

type WebHandler struct {
	svc        *service.URLService
	generateID func() string
	tmpl       *template.Template
}

func NewWebHandler(svc *service.URLService, generateID func() string) *WebHandler {
	tmpl := template.Must(template.ParseFiles("/home/nexus/Desktop/my-project/pet/url-shortener/src/internal/web/templates/index.html"))
	return &WebHandler{svc: svc, generateID: generateID, tmpl: tmpl}
}

func (h *WebHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.tmpl.Execute(w, nil)
}

func (h *WebHandler) ShortenWeb(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	id, err := h.svc.Shorten(url, h.generateID)
	if err != nil {
		http.Error(w, "failed to shorten url", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"ShortURL": "http://localhost:8080/" + id,
	}
	h.tmpl.Execute(w, data)
}
