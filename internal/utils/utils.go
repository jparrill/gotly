package utils

import (
	"os"
	"path/filepath"

	gotlyTypes "github.com/jparrill/gotly/internal/types"
	"gopkg.in/yaml.v2"
)

func GetBaseRepoPath(fl string) (string, error) {
	if _, err := os.Stat(fl + "/assets"); err != nil {
		return "", os.ErrNotExist
	}

	return filepath.Abs(fl)
}

func LoadSetPathFromYAML(yml []byte) (map[string]string, error) {
	var urlMap []gotlyTypes.URLMap
	urlPaths := make(map[string]string)

	err := yaml.Unmarshal(yml, &urlMap)
	if err != nil {
		return nil, err
	}
	for _, set := range urlMap {
		urlPaths[set.Path] = set.Url
	}

	return urlPaths, nil
}
