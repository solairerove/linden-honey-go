package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/solairerove/linden-honey-go/model"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	lindenHoneyScraperURL   = "LINDEN_HONEY_SCRAPER_URL"
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
		log.Fatal("Error loading .env file", err)
	}

	b, err := strconv.ParseBool(os.Getenv(usingLindenHoneyScraper))
	if err != nil {
		log.Fatal("Can't parse to bool value: ", err)
	}

	log.Println(usingLindenHoneyScraper, b)

	scraperURL := fmt.Sprintf("%s/songs", os.Getenv(lindenHoneyScraperURL))
	log.Println(lindenHoneyScraperURL, scraperURL)

	resp, err := http.Get(scraperURL)
	if err != nil {
		log.Fatal("Error loading data from scraper: ", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing body from scraper: ", err)
		}
	}()

	songs := make([]model.Song, 0)

	err = json.NewDecoder(resp.Body).Decode(&songs)
	if err != nil {
		log.Fatal("Error unmarshalling data from scraper: ", err)
	}

	log.Println(len(songs))
}
