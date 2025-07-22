package main

import (
	"fmt"
	"log"
	"urlshortener/internal/shortener"
	"urlshortener/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	svc := shortener.NewService(store)

	longURL := "https://example.com/very/long/url"
	shortKey, err := svc.Shorten(longURL)
	if err != nil {
		log.Fatalf("failed to shorten: %v", err)
	}

	fmt.Printf("Short URL: %s\n", shortKey)

	originalURL, err := svc.Resolve(shortKey)
	if err != nil {
		log.Fatalf("failed to resolve: %v", err)
	}
	fmt.Printf("Resolved URL: %s\n", originalURL)
}
