package service

import (
	"awura-shortener/internal/model"
	"awura-shortener/internal/repository"
	"errors"
	"math/rand"
	"time"
)

const (
	shortURLLength = 5
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type URLService struct {
	Repo *repository.URLRepository
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
	if originalURL == "" {
		return "", errors.New("original URL cannot be empty")
	}

	// Generate a unique short URL
	shortURL := s.generateShortURL()

	// Ensure the short URL is unique by checking the database
	for {
		exists, err := s.Repo.ShortURLExists(shortURL)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
		shortURL = s.generateShortURL()
	}

	// Save the short URL and original URL in the database
	url := &model.URL{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	}
	err := s.Repo.CreateURL(url)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *URLService) GetOriginalURL(shortURL string) (string, error) {
	if shortURL == "" {
		return "", errors.New("short URL cannot be empty")
	}

	// Retrieve the original URL from the database
	url, err := s.Repo.GetURL(shortURL)
	if err != nil {
		return "", err
	}
	if url == nil {
		return "", errors.New("short URL not found")
	}

	return url.OriginalURL, nil
}

func (s *URLService) generateShortURL() string {
	b := make([]byte, shortURLLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
