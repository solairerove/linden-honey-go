package model

// Song ... tbd
type Song struct {
	Title  string        `json:"title,omitempty"`
	Link   string        `json:"link,omitempty"`
	Author string        `json:"author,omitempty"`
	Album  string        `json:"album,omitempty"`
	Verses []Verse       `json:"verses,omitempty"`
}

// Verse ... tbd
type Verse struct {
	Order  int           `json:"order"`
	Verse  string        `json:"verse,omitempty"`
}
