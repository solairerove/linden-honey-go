package example

import (
	"github.com/solairerove/linden-honey-go/model"
	"github.com/solairerove/linden-honey-go/sarvar"
)

func createSong() {
	s := sarvar.Sarvar{}

	// connect to db with credentials, os.env variables in non local machine
	s.Initialize("linden-honey-user", "linden-honey-pass", "linden-honey")

	// test data to improve save method
	verses := []model.Verse{model.Verse{Ordinal: 1, Data: "lyrics"}, model.Verse{Ordinal: 2, Data: "lyrics"}}
	song := model.Song{
		Title:  "Flying Home",
		Link:   "http://www.gr-oborona.ru/texts/1056561331.html",
		Author: "американская народная, поёт Д. Селиванов",
		Album:  "Хроника Пикирующего Бомбардировщика",
		Verses: verses}

	song.SaveSong(s.DB)
}
