package sarvar

import (
	"database/sql"
	"fmt"
	"log"
)

// Sarvar ... tbd
type Sarvar struct {
	DB *sql.DB
}

// Initialize ... tbd
func (s *Sarvar) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	s.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
}
