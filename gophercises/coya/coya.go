package coya

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type story struct {
	Story    map[string]chapter
	template *template.Template
}

type option struct {
	Text        string
	ChapterName string `json:"arc"`
}

type chapter struct {
	Title   string
	Story   []string
	Options []option
}

// NewStoryHandler returns an http.Handler for the new story
func NewStoryHandler(jsonFile, templateFile string) http.Handler {
	s := story{
		Story:    make(map[string]chapter),
		template: template.Must(template.ParseFiles(templateFile)),
	}
	s.loadStory(jsonFile)
	return s.httpHandler()
}

func (s *story) loadStory(jsonFile string) {
	file, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(file)
	if err = d.Decode(&s.Story); err != nil {
		panic(err)
	}
	log.Println("Story successfully decoded")
}

func (s *story) httpHandler() http.Handler {
	mux := http.NewServeMux()
	for chapterName := range s.Story {
		chapterName := chapterName
		route := fmt.Sprintf("/%s", chapterName)
		log.Println("Registering http handler for", route)
		mux.HandleFunc(route, func(res http.ResponseWriter, req *http.Request) {
			s.responseWriter(res, req, chapterName)
		})
	}
	return mux
}

func (s *story) responseWriter(res http.ResponseWriter, req *http.Request, chapterName string) {
	s.template.Execute(res, s.Story[chapterName])
}
