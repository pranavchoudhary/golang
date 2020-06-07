package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type article struct {
	Id          string `json:id`
	Title       string `json:title`
	Description string `json:description`
	Content     string `json:content`
}

var articles map[string]article = map[string]article{
	"lotr1": {
		Id:          "lotr1",
		Title:       "Fellowship of the Ring",
		Description: "First book in LOTR trilogy",
		Content:     "Story of creation of the fellowship",
	},
	"lotr2": {
		Id:          "lotr2",
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

func returnOneArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint hit: articles/", key)
	json.NewEncoder(w).Encode(articles[key])
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var a article
	if err := json.Unmarshal(reqBody, &a); err != nil {
		fmt.Println("Endpoint hit: [POST] articles/", err)
		return
	}
	fmt.Println("Endpoint hit: [POST] articles/", a.Id)
	articles[a.Id] = a
	json.NewEncoder(w).Encode(a)
}

func setupEndpoints() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", returnOneArticle)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	fmt.Println("Starting v2 of REST API server")
	setupEndpoints()
}
