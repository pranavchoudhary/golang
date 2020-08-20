package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pranavchoudhary/golang/gophercises/urlshort"
)

type staticURLSource struct{}

var redirectMap = map[string]string{
	"/go":   "https://golang.org",
	"/anet": "https://www.arista.com/",
	"/lwn":  "https://lwn.net/",
}

func (s staticURLSource) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if redirectURL, ok := redirectMap[req.RequestURI]; ok {
		http.Redirect(res, req, redirectURL, http.StatusFound)
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, "Unknown short URL")
	}
}

// yamlFile and jsonFile are set from command line arguments
var (
	yamlFile string
	jsonFile string
)

func init() {
	flag.StringVar(&yamlFile, "yaml", "", "Yaml file to read the short urls from")
	flag.StringVar(&jsonFile, "json", "", "JSON file to read the short URLS from")
}

func main() {
	var handler http.Handler
	flag.Parse()
	if yamlFile != "" {
		handler = urlshort.NewYamlHandler(yamlFile)
	} else if jsonFile != "" {
		handler = urlshort.NewJsonHandler(jsonFile)
	} else {
		handler = staticURLSource{}
	}
	log.Fatalln(http.ListenAndServe(":8081", handler))
}
