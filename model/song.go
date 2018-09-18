package model

import (
	"database/sql"
	"log"

	"github.com/satori/go.uuid"
)

// Song ... tbd
type Song struct {
	ID     uuid.NullUUID `sql:",pk,type:uuid default uuid_generate_v4()" json:"Id"`
	Title  string        `json:"title,omitempty"`
	Link   string        `json:"link,omitempty"`
	Author string        `json:"author,omitempty"`
	Album  string        `json:"album,omitempty"`
	Verses *[]Verse      `json:"verses,omitempty"`
}

// Verse ... tbd
type Verse struct {
	ID      uuid.NullUUID `sql:",pk,type:uuid default uuid_generate_v4()" json:"Id"`
	Ordinal int           `json:"ord"`
	Data    string        `json:"data,omitempty"`
	SongID  uuid.NullUUID `sql:",type:uuid" json:"songId"`
}

// CreateSong ... tbd
func (s *Song) CreateSong(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO songs(title, link, author, album) VALUES($1, $2, $3, $4) RETURNING id",
		s.Title, s.Link, s.Author, s.Album).Scan(&s.ID)

	if err != nil {
		return err
	}

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
