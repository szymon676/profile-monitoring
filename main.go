package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{username}", handleUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":1337", handlers.CORS()(router)))
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	username := mux.Vars(r)["username"]
	i := 0

	var followers string
	var following string
	c.OnHTML("span.text-bold.color-fg-default", func(e *colly.HTMLElement) {
		if i == 0 {
			followers = e.Text
			i++
		} else if i == 1 {
			following = e.Text
		}
	})

	url := fmt.Sprintf("https://github.com/%s", username)
	c.Visit(url)
	returnString := fmt.Sprintf("Followers: %v Following: %v", followers, following)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(returnString)
}
