package urlshort

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jparrill/gotly/internal/utils"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
		log.Printf("Calling path %s", req.URL.Path)
		if url != "" {
			http.Redirect(w, req, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, req)
	}
}

func YAMLHandler(yml []byte, ctx context.Context, rdb *redis.Client, fallback http.Handler) (http.HandlerFunc, error) {
	// Grab entries from YAML file
	urlPaths := make(map[string]string)

	err := utils.LoadSetPathFromYAML(yml, &urlPaths)
	if err != nil {
		log.Println("WARNING: Assets YAML File Is not accesible")
	}

	// Grab entries from DDBB
	err = utils.LoadKVFromRedis(ctx, rdb, &urlPaths)
	if err != nil {
		log.Println("WARNING: REDIS DDBB Is not accesible")
	}

	return MapHandler(urlPaths, fallback), nil
}
