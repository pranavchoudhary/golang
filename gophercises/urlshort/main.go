package main

import (
	"net/http"
)

var db map[string]string = map[string]string{
	"ggl":  "http://www.google.com",
	"anet": "http://www.arista.com",
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]
	if url, ok := db[shortUrl]; ok {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8081", nil)
}
