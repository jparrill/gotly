package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	gotlyDB "github.com/jparrill/gotly/internal/database"
	"github.com/jparrill/gotly/pkg/urlshort"
)

func Run(ctx context.Context, basePath string) {
	mux := defaultMux()

	// DDBB
	// - TODO: Cobra + Config file to populate DDBB options
	rdOpts := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// Create RDB Client
	rdb := gotlyDB.GetClient(ctx, &rdOpts)

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

	yamlHandler, err := urlshort.YAMLHandler([]byte(ymlB), ctx, &rdb, mapHandler)
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
