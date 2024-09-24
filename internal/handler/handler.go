package handler

import (
	"awura-shortener/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type URLHandler struct {
	Service *service.URLService
}

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := h.Service.ShortenURL(req.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	baseURL := fmt.Sprintf("%s://%s", "http", r.Host)

	fullShortURL := fmt.Sprintf("%s/r/%s", baseURL, shortURL)

	resp := ShortenResponse{ShortURL: fullShortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/r/"):]
	originalURL, err := h.Service.GetOriginalURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
