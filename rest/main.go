package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type article struct {
	Title       string `json:title`
	Description string `json:description`
	Content     string `json:content`
}

var articles []article = []article{
	{
		Title:       "Fellowship of the Ring",
		Description: "First book in LOTR trilogy",
		Content:     "Story of creation of the fellowship",
	},
	{
		Title:       "The Two Towers",
		Description: "Second book in the LOTR trilogy",
		Content:     "Story of Isengard and Rohan",
	},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page!")
	fmt.Println("Endpoint hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: articles")
	json.NewEncoder(w).Encode(articles)
}

func setupEndpoints() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	setupEndpoints()
}
