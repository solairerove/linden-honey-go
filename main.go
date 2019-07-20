package main

import (
	_ "github.com/lib/pq"
	sarvar "github.com/solairerove/linden-honey-go/sarvar"
)

const (
	dbUsername = "linden-honey-user"
	dbPassword = "linden-honey-pass"
	dbName     = "linden-honey"
	dbPort     = "5430"
)

func main() {
	s := sarvar.Sarvar{}

	// connect to db with credentials, os.env variables in non local machine
	s.Initialize(dbUsername, dbPassword, dbName, dbPort)

	s.Run(":8000")
}
