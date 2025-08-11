package rredis

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type closerFunc func()

// NewClientFromViper constructs a go-redis client using global flags/env.
// Env vars (examples): RHM_ADDR, RHM_USER, RHM_PASSWORD, RHM_DB, RHM_TLS, RHM_TLS_SKIP_VERIFY
func NewClientFromViper(ctx context.Context) (*redis.Client, closerFunc, error) {
	addr := viper.GetString("addr")
	user := viper.GetString("user")
	pass := viper.GetString("password")
	db := viper.GetInt("db")
	useTLS := viper.GetBool("tls")
	insecure := viper.GetBool("tls_skip_verify")

	var tlsCfg *tls.Config
	if useTLS {
		tlsCfg = &tls.Config{
			InsecureSkipVerify: insecure, // user-controlled
		}
	}

	opts := &redis.Options{
		Addr:      addr,
		Username:  user,
		Password:  pass,
		DB:        db,
		TLSConfig: tlsCfg,
		// Let commands honor context timeouts
	}

	client := redis.NewClient(opts)

	// Quick sanity ping with the root timeout (or 3s fallback)
	timeout := viper.GetDuration("timeout")
	if timeout <= 0 {
		timeout = 3 * time.Second
	}
	pingCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return client, func() { _ = client.Close() }, nil
}

// PreflightNonReplica returns error if server is a replica (optional to use later).
func PreflightNonReplica(ctx context.Context, c *redis.Client) error {
	res := c.Info(ctx, "replication")
	if res.Err() != nil {
		return res.Err()
	}
	text := res.Val()
	// cheap/fast check
	if contains(text, "\nrole:slave") || contains(text, "\nrole:replica") {
		return errors.New("connected to a replica; some operations may be inaccurate")
	}
	return nil
}

// tiny helper to avoid importing strings just for this
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (len(sub) == 0 || indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	// naive search (small strings only here)
outer:
	for i := 0; i+len(sub) <= len(s); i++ {
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
