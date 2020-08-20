package urlshort

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type jsonURLDb struct {
	urlMap map[string]string
}

// NewJsonHandler returns an http.Handler using json as database
func NewJsonHandler(filename string) http.Handler {
	ret := jsonURLDb{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicln("Failed to read file", filename, err)
	}
	err = json.Unmarshal(data, &ret.urlMap)
	if err != nil {
		log.Panicln(err)
	}
	return &ret
}

func (j *jsonURLDb) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("Request received for", req.RequestURI)
	if redirectURL, ok := j.urlMap[req.RequestURI]; ok {
		http.Redirect(res, req, redirectURL, http.StatusFound)
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, "Unknown URI", req.RequestURI)
	}
}
