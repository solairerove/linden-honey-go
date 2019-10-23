package model

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"log"
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
	Ordinal int           `json:"order"`
	Verse   string        `json:"verse,omitempty"`
	SongID  uuid.NullUUID `sql:",type:uuid" json:"-"`
}

// SaveSong ... tbd
func (s *Song) SaveSong(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO songs(title, link, author, album) VALUES($1, $2, $3, $4) RETURNING id",
		s.Title, s.Link, s.Author, s.Album).Scan(&s.ID)

	if err != nil {
		return err
	}

	log.Printf("Persisted song id -> %s", s.ID.UUID.String())

	for _, v := range s.Verses {
		v.SongID = s.ID

		v.saveVerse(db)
	}

	return nil
}

// saveVerse ... tbd
func (v *Verse) saveVerse(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO verses(ordinal, verse, song_id) VALUES($1, $2, $3) RETURNING id",
		v.Ordinal, v.Verse, v.SongID).Scan(&v.ID)

	if err != nil {
		return err
	}

	log.Printf("Persisted verse id -> %s", v.ID.UUID.String())

	return nil
}

// FindSongByID ... tbd
func FindSongByID(db *sql.DB, id string) (Song, error) {

	rows, err := db.Query(`
	SELECT songs.id, songs.title, songs.link, songs.author, songs.album, verses.ordinal, verses.verse
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
		if err := rows.Scan(&s.ID, &s.Title, &s.Link, &s.Author, &s.Album, &v.Ordinal, &v.Verse); err != nil {
			log.Fatalf("Such error: %s", err.Error())
			return Song{}, err
		}
		verses = append(verses, v)
	}

	s.Verses = verses

	return s, nil
}

// FetchNameToIDMapByName ... tbd
func FetchNameToIDMapByName(db *sql.DB, name string) (map[string]string, error) {
	rows, err := db.Query(`SELECT songs.id, songs.title FROM songs WHERE songs.title ILIKE '%` + name + `%'`)

	if err != nil {
		log.Fatalf("Such SQL error: %s", err.Error())
		return make(map[string]string, 0), err
	}

	defer rows.Close()

	nameToID := make(map[string]string, 0)

	for rows.Next() {
		var s Song
		if err := rows.Scan(&s.ID, &s.Title); err != nil {
			log.Fatalf("Such rows mapping error: %s", err.Error())
			return make(map[string]string, 0), err
		}

		nameToID[s.ID.UUID.String()] = s.Title
	}

	return nameToID, nil
}
