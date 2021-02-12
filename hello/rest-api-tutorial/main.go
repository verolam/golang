package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article

type Article struct {
	Title   string "json:'Title'"
	Desc    string "json:'desc'"
	Content string "json:'content'"
}

// Articles : array of Article
var Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		Article{Title: "Test Title", Desc: "Description", Content: "Content"},
	}

	fmt.Println("Endpoint Hit : all Articles Endpoint")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
