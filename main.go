package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/solairerove/linden-honey-go/model"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	lindenHoneyScraperURL   = "LINDEN_HONEY_SCRAPER_URL"
	usingLindenHoneyScraper = "USING_LINDEN_HONEY_SCRAPER"

	// TODO: to .env
	dbUsername = "linden-honey-user"
	dbPassword = "linden-honey-pass"
	dbName     = "linden-honey"
)

var db *sql.DB

func main() {
	log.Println("Starting application")

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUsername, dbPassword, dbName, "5432")

	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	// get scraper bool from .env file
	useScrapper, err := strconv.ParseBool(os.Getenv(usingLindenHoneyScraper))
	if err != nil {
		log.Fatal("Can't parse to bool value: ", err)
	}

	log.Println(usingLindenHoneyScraper, useScrapper)

	scraperURL := fmt.Sprintf("%s/songs", os.Getenv(lindenHoneyScraperURL))
	log.Println(lindenHoneyScraperURL, scraperURL)

	// TODO: rewrite and move to package after go web udemy
	// fetch and update collection if necessary
	if useScrapper {
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

		for _, s := range songs {
			s.SaveSong(db)
		}

		log.Println(songs[0])
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", handlers.CompressHandler(router)))
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	if err != nil {
		log.Fatalf("Something wrong with greeting endppoint: %v", err)
	}
}
