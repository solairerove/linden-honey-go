package example

import (
	"encoding/json"
	"log"

	"github.com/solairerove/linden-honey-go/model"
	"github.com/solairerove/linden-honey-go/sarvar"
)

func findByName() {
	s := sarvar.Sarvar{}

	// connect to db with credentials, os.env variables in non local machine
	s.Initialize("linden-honey-user", "linden-honey-pass", "linden-honey", "5430")

	unmarshaledMap, _ := model.FetchNameToIDMapByName(s.DB, "м")

	marshaledMap, _ := json.Marshal(unmarshaledMap)
	log.Printf("Find current map %s", string(marshaledMap))
}
