package main

import (
	"awura-shortener/internal/handler"
	"awura-shortener/internal/repository"
	"awura-shortener/internal/service"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createTable(db); err != nil {
		log.Fatal("Failed to create table:", err)
	}

	repo := &repository.URLRepository{DB: db}
	srv := &service.URLService{Repo: repo}
	h := &handler.URLHandler{Service: srv}

	http.HandleFunc("/shorten", h.CreateShortURL)
	http.HandleFunc("/r/", h.Redirect)

	log.Println("Starting server on :18080")
	log.Fatal(http.ListenAndServe(":18080", nil))
}

func createTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        short_url VARCHAR(255) NOT NULL UNIQUE,
        original_url TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    )`
	_, err := db.Exec(query)
	return err
}
