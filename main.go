package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/solairerove/linden-honey-go/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// TODO: rewrite and move to package after go web udemy
	// new mongodb client
	var auth options.Credential
	auth.Username = dbUsername
	auth.Password = dbPassword

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(auth))

	if err != nil {
		log.Fatal("Error connecting to mongodb: ", err)
	}

	// connect to context?
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("What the fuck is context todo: ", err)
	}

	// check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("What the fuck is context todo ping: ", err)
	}

	// TODO: to .env
	collection := client.Database("linden_honey").Collection("songs")

	// close connection
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Error closing mongo client: ", err)
		}
	}()
	log.Println("Connection to MongoDB closed.")

	// get scraper bool from .env file
	b, err := strconv.ParseBool(os.Getenv(usingLindenHoneyScraper))
	if err != nil {
		log.Fatal("Can't parse to bool value: ", err)
	}

	log.Println(usingLindenHoneyScraper, b)

	scraperURL := fmt.Sprintf("%s/songs", os.Getenv(lindenHoneyScraperURL))
	log.Println(lindenHoneyScraperURL, scraperURL)

	// TODO: rewrite and move to package after go web udemy
	// fetch and update collection if necessary
	if b {
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

		// persist collection plz
		for _, song := range songs {
			insertResult, err := collection.InsertOne(context.Background(), song)
			if err != nil {
				log.Fatal("Error persisting songs into mongodb: ", err)
			}

			fmt.Println("Inserted one document: ", insertResult.InsertedID)
		}

		log.Println(len(songs))
	}
}
