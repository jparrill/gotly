package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func GetClient(ctx context.Context, dbOpt *redis.Options) redis.Client {
	return *redis.NewClient(dbOpt)
}
