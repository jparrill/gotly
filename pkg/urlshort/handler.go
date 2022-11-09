package urlshort

import (
	"net/http"

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

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlPaths, err := utils.LoadSetPathFromYAML(yml)
	if err != nil {
		return nil, err
	}

	return MapHandler(urlPaths, fallback), nil
}
