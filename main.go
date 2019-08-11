package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const (
	usingLindenHoneyScraper = "USING_LINDEN_HONEY_SCRAPER"
	dbUsername              = "linden-honey-user"
	dbPassword              = "linden-honey-pass"
	dbName                  = "linden-honey"
	dbPort                  = "5430"
)

func main() {
	log.Println("Nothing else matters")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	b, err := strconv.ParseBool(os.Getenv(usingLindenHoneyScraper))
	if err != nil {
		log.Fatal("Can't parse to bool value")
	}

	log.Println(usingLindenHoneyScraper, b)
}
