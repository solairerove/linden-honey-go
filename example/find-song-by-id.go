package example

import (
	"encoding/json"
	"log"

	"github.com/solairerove/linden-honey-go/model"
	"github.com/solairerove/linden-honey-go/sarvar"
)

var s = sarvar.Sarvar{}

func findByID() {
	s := sarvar.Sarvar{}

	// connect to db with credentials, os.env variables in non local machine
	s.Initialize("linden-honey-user", "linden-honey-pass", "linden-honey")

	unmarshaledSong, _ := model.FindSongByID(s.DB, "08b40777-c464-4063-9ba7-418566c09bac")

	marshaledSong, _ := json.Marshal(unmarshaledSong)
	log.Printf("Find current song %s", string(marshaledSong))
}
