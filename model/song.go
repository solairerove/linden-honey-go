package model

import (
	uuid "github.com/satori/go.uuid"
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
	ID     uuid.NullUUID `sql:",pk,type:uuid default uuid_generate_v4()" json:"-"`
	Order  int           `json:"order"`
	Verse  string        `json:"verse,omitempty"`
	SongID uuid.NullUUID `sql:",type:uuid" json:"-"`
}
