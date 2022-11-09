package utils

import (
	"context"
	"os"
	"path/filepath"

	"github.com/go-redis/redis/v8"
	gotlyTypes "github.com/jparrill/gotly/internal/types"
	"gopkg.in/yaml.v2"
)

func GetBaseRepoPath(fl string) (string, error) {
	if _, err := os.Stat(fl + "/assets"); err != nil {
		return "", os.ErrNotExist
	}

	return filepath.Abs(fl)
}

func LoadSetPathFromYAML(yml []byte, tuples *map[string]string) error {
	var urlMap []gotlyTypes.URLMap

	err := yaml.Unmarshal(yml, &urlMap)
	if err != nil {
		return err
	}
	for _, set := range urlMap {
		(*tuples)[set.Path] = set.Url
	}

	return nil
}

func LoadKVFromRedis(ctx context.Context, rdb *redis.Client, tuples *map[string]string) error {

	iter := rdb.Scan(ctx, 0, "/*", 0).Iterator()
	for iter.Next(ctx) {
		get := rdb.Get(ctx, iter.Val())
		(*tuples)[iter.Val()] = get.Val()
	}
	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
