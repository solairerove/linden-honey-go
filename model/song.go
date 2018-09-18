package model

import (
	"database/sql"
	"log"
)

// Song ... tbd
type Song struct {
	Title  string   `json:"title,omitempty"`
	Link   string   `json:"link,omitempty"`
	Author string   `json:"author,omitempty"`
	Album  string   `json:"album,omitempty"`
	Verses *[]Verse `json:"verses,omitempty"`
}

// Verse ... tbd
type Verse struct {
	Ordinal int    `json:"ord"`
	Data    string `json:"data,omitempty"`
}

// CreateSong ... tbd
func (s *Song) CreateSong(db *sql.DB) error {
	log.Fatal("Not implemented yet")
	return nil
}

// GetSong ... tbd
func (s *Song) GetSong(db *sql.DB) error {
	log.Fatal("Not implemented yet")
	return nil
}

// GetSongs ... tbd
func GetSongs(db *sql.DB, start, count int) ([]Song, error) {
	log.Fatal("Not implemented yet")
	return nil, nil
}
