package rredis

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
)

// Example preflight helpers (to be expanded Day 15).
func IsCluster(ctx context.Context, cli redis.Cmdable) (bool, error) {
	info, err := cli.Info(ctx, "cluster").Result()
	if err != nil {
		return false, err
	}
	return strings.Contains(info, "cluster_enabled:1"), nil
}
