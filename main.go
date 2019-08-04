package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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

	log.Println(usingLindenHoneyScraper, os.Getenv(usingLindenHoneyScraper))
}
