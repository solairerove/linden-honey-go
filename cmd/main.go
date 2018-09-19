package main

import (
	"encoding/json"
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

	/*
		// test data to improve save method
		verses := []model.Verse{model.Verse{Ordinal: 1, Data: "lyrics"}, model.Verse{Ordinal: 2, Data: "lyrics"}}
		song := model.Song{
			Title:  "Flying Home",
			Link:   "http://www.gr-oborona.ru/texts/1056561331.html",
			Author: "американская народная, поёт Д. Селиванов",
			Album:  "Хроника Пикирующего Бомбардировщика",
			Verses: verses}

		song.CreateSong(s.DB)
	*/

	unmarshaledSong, _ := model.FindSongByID(s.DB, "08b40777-c464-4063-9ba7-418566c09bac")

	marshaledSong, _ := json.Marshal(unmarshaledSong)
	log.Printf("Find current song %s", string(marshaledSong))

	unmarshaledMap, _ := model.FetchNameToIDMapByName(s.DB, "м")

	marshaledMap, _ := json.Marshal(unmarshaledMap)
	log.Printf("Find current map %s", string(marshaledMap))
}

// future migration
func ensureTableExists() {
	if _, err := s.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// future migration pff
const tableCreationQuery = ``
