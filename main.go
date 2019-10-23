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

type Server struct {
	db *sql.DB
}

func main() {
	log.Println("Starting application")

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	server := Server{}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUsername, dbPassword, dbName, "5432")

	server.db, err = sql.Open("postgres", connectionString)

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
			s.SaveSong(server.db)
		}

		log.Println(songs[0])
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/songs/{id}", server.getSongByID).Methods("GET")
	router.HandleFunc("/songs", server.getSongByName).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", handlers.CompressHandler(router)))
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	if err != nil {
		log.Fatalf("Something wrong with greeting endppoint: %v", err)
	}
}

func (s *Server) getSongByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	song, err := model.FindSongByID(s.db, id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, song)
}

func (s *Server) getSongByName(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	names, ok := vals["name"]
	var name string
	if ok {
		if len(names) >= 1 {
			name = names[0]
		}
	}

	songsMap, err := model.FetchNameToIDMapByName(s.db, name)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, songsMap)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
