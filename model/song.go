package model

import (
	"database/sql"
	"log"

	"github.com/satori/go.uuid"
)

// Song ... tbd
type Song struct {
	ID     uuid.NullUUID `sql:",pk,type:uuid default uuid_generate_v4()" json:"-"`
	Title  string        `json:"title,omitempty"`
	Link   string        `json:"link,omitempty"`
	Author string        `json:"author,omitempty"`
	Album  string        `json:"album,omitempty"`
	Verses []Verse       `json:"verses,omitempty"`
}

// Verse ... tbd
type Verse struct {
	ID      uuid.NullUUID `sql:",pk,type:uuid default uuid_generate_v4()" json:"-"`
	Ordinal int           `json:"ord"`
	Data    string        `json:"data,omitempty"`
	SongID  uuid.NullUUID `sql:",type:uuid" json:"-"`
}

// CreateSong ... tbd
func (s *Song) CreateSong(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO songs(title, link, author, album) VALUES($1, $2, $3, $4) RETURNING id",
		s.Title, s.Link, s.Author, s.Album).Scan(&s.ID)

	if err != nil {
		return err
	}

	log.Printf("Persisted song id -> %s", s.ID.UUID.String())

	for _, v := range s.Verses {
		v.SongID = s.ID

		v.CreateVerse(db)
	}

	return nil
}

// CreateVerse ... tbd
func (v *Verse) CreateVerse(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO verses(ordinal, data, song_id) VALUES($1, $2, $3) RETURNING id",
		v.Ordinal, v.Data, v.SongID).Scan(&v.ID)

	if err != nil {
		return err
	}

	log.Printf("Persisted verse id -> %s", v.ID.UUID.String())

	return nil
}

// GetSong ... tbd
func GetSong(db *sql.DB, id string) (Song, error) {

	rows, err := db.Query(`
	SELECT songs.id, songs.title, songs.link, songs.author, songs.album, verses.id, verses.ordinal, verses.data, verses.song_id
	FROM songs 
		INNER JOIN verses 
		ON songs.id=verses.song_id 
	WHERE songs.id = $1`, id)

	if err != nil {
		log.Fatalf("Such error: %s", err.Error())
		return Song{}, err
	}

	defer rows.Close()

	var s Song
	verses := []Verse{}

	for rows.Next() {

		var v Verse
		if err := rows.Scan(&s.ID, &s.Title, &s.Link, &s.Author, &s.Album, &v.ID, &v.Ordinal, &v.Data, &v.SongID); err != nil {
			log.Fatalf("Such error: %s", err.Error())
			return Song{}, err
		}
		verses = append(verses, v)
	}

	s.Verses = verses

	return s, nil
}

// GetSongs ... tbd
func GetSongs(db *sql.DB, start, count int) ([]Song, error) {
	log.Fatal("Not implemented yet")
	return nil, nil
}
