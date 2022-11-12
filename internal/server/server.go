package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jparrill/gotly/internal/config"
	gotlyDB "github.com/jparrill/gotly/internal/database"
	"github.com/jparrill/gotly/pkg/urlshort"
)

func Run(ctx context.Context, basePath string, config *config.Config) {
	// Create RDB Client
	rdb := gotlyDB.GetClient(ctx, &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.DB.Host, config.DB.Port),
		Password: config.DB.Pass,
		DB:       config.DB.Num,
	})

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/go":   "https://google.es",
		"/fino": "https://finofilipino.org",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, http.NewServeMux())

	// Load additional sample file located in assets
	ymlB, err := os.ReadFile(config.SourceFile)
	if err != nil {
		log.Fatalf("Error loading data from source file %s: %v", config.SourceFile, err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(ymlB), ctx, &rdb, mapHandler)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Starting the server on %d port", config.AppPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.AppPort), yamlHandler)
	if err != nil {
		log.Panicf("Failed to raise up webserver in %d port", config.AppPort)
	}
}
