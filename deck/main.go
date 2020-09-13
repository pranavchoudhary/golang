package main

import (
	"flag"
)

var (
	yamlFile *string
	jsonFile *string
)

func init() {
	yamlFile = flag.String("yamlFile", "", "YAML file to read the shortened URLs from")
	jsonFile = flag.String("jsonFile", "", "JSON file to read the shortened URLs from")
}

func main() {
	flag.Parse()
	startServer()
}
