package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/solairerove/linden-honey-go/model"
	"github.com/solairerove/linden-honey-go/sarvar"
)

var s = sarvar.Sarvar{}

func main() {
	s := sarvar.Sarvar{}

	// connect to db with credentials, os.env variables in non local machine
	s.Initialize("linden-honey-user", "linden-honey-pass", "linden-honey")

	// my super migration
	// ensureTableExists()

	// test data to improve save method
	verses := []model.Verse{model.Verse{Ordinal: 1, Data: "lyrics"}, model.Verse{Ordinal: 2, Data: "lyrics"}}
	song := model.Song{
		Title:  "Flying Home",
		Link:   "http://www.gr-oborona.ru/texts/1056561331.html",
		Author: "американская народная, поёт Д. Селиванов",
		Album:  "Хроника Пикирующего Бомбардировщика",
		Verses: verses}

	song.CreateSong(s.DB)
}

// future migration
func ensureTableExists() {
	if _, err := s.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// future migration pff
const tableCreationQuery = ``
