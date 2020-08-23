package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pranavchoudhary/golang/gophercises/coya"
)

var (
	listenAddr   string
	jsonFile     string
	templateFile string
)

func init() {
	flag.StringVar(&listenAddr, "addr", ":8081", "Address and port number on which the http server shall listen")
	flag.StringVar(&jsonFile, "file", "", "Json file containing the story")
	flag.StringVar(&templateFile, "template", "", "HTML template file")
}

func main() {
	flag.Parse()
	log.Panicln(http.ListenAndServe(listenAddr, coya.NewStoryHandler(jsonFile, templateFile)))
}
