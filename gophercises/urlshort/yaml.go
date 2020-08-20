package urlshort

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlURLgSource struct {
	urlMap map[string]string
}

// NewYamlHandler creates and returns a http.Handler service redirects from
// the given yaml file
func NewYamlHandler(filename string) http.Handler {
	data, err := ioutil.ReadFile(filename)
	ret := &yamlURLgSource{
		urlMap: map[string]string{},
	}
	if err != nil {
		log.Fatalln("Failed to read yaml file", err)
		return ret
	}
	if err = yaml.Unmarshal(data, &ret.urlMap); err != nil {
		log.Fatalln("Failed to Unmarshal YAML file", err)
	}
	log.Println("Short URLS:", ret)
	return ret
}

func (y *yamlURLgSource) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("Request received for", req.RequestURI)
	if redirectURL, ok := y.urlMap[req.RequestURI]; ok {
		http.Redirect(res, req, redirectURL, http.StatusFound)
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, "Unknown URI", req.RequestURI)
	}
}
