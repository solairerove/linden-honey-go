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
