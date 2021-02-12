package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

type Article struct {
	Title   string "json:'Title'"
	Desc    string "json:'desc'"
	Content string "json:'content'"
}

var Articles []Article

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit : all Articles Endpoint")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Description", Content: "Content"},
		Article{Title: "Hello 2", Desc: "Description", Content: "Content"},
	}

	handleRequests()

}
