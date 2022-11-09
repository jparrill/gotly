package urlshort

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jparrill/gotly/internal/utils"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
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
		return nil, err
	}

	// Grab entries from DDBB
	err = utils.LoadKVFromRedis(ctx, rdb, &urlPaths)
	if err != nil {
		return nil, err
	}

	return MapHandler(urlPaths, fallback), nil
}
