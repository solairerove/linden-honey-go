package main

import (
	"log"

	"github.com/solairerove/linden-honey-go/sarvar"
)

var s = sarvar.Sarvar{}

func main() {
	s := sarvar.Sarvar{}

	// connect to db with credentials, os.env variables in non local machine
	s.Initialize("linden-honey-user", "linden-honey-pass", "linden-honey")

	// my super migration
	ensureTableExists()
}

func ensureTableExists() {
	if _, err := s.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = ``
