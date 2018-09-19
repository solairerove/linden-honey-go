package sarvar

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/solairerove/linden-honey-go/model"
)

// Sarvar ... tbd
type Sarvar struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize ... tbd
func (s *Sarvar) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	s.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

// Run ... tbd
func (s *Sarvar) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

func (s *Sarvar) initializeRoutes() {
	s.Router.HandleFunc("/hello", s.helloHandle).Methods("GET")
	s.Router.HandleFunc("/songs/{id}", s.getSongByID).Methods("GET")
	s.Router.HandleFunc("/songs", s.getSongByName).Methods("GET")
}

func (s *Sarvar) helloHandle(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Hello World !")
}

func (s *Sarvar) getSongByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	song, err := model.FindSongByID(s.DB, id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, song)
}

func (s *Sarvar) getSongByName(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	names, ok := vals["name"]
	var name string
	if ok {
		if len(names) >= 1 {
			name = names[0]
		}
	}

	songsMap, err := model.FetchNameToIDMapByName(s.DB, name)

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
