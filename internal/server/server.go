package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jparrill/gotly/pkg/urlshort"
)

func Run(basePath string) {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/go":   "https://google.es",
		"/fino": "https://finofilipino.org",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Load additional sample file located in assets
	ymlB, err := ioutil.ReadFile(basePath + "/assets/samples/additionalSets.yaml")
	if err != nil {
		log.Fatalf("Error loading assets sample file: %v", err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(ymlB), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
