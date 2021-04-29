package repository

import (
	"context"
	"time"
)

type KVIface interface {
	Incr(ctx context.Context, key string) error
	IncrAndGet(ctx context.Context, key string) (int64, error)
	Set(ctx context.Context, key string, value interface{}) error
	SetWithTTL(ctx context.Context, key string, value interface{}) error
	Del(ctx context.Context, key string) (bool, error)
	MSet(ctx context.Context, pair ...Pair) error
	MSetWithTTL(ctx context.Context, expiration time.Duration, pair ...Pair) error
	HIncr(ctx context.Context, key, field string, incr int64) error
	Exists(ctx context.Context, key string) (bool, error)
	MGet(ctx context.Context, keys ...string) ([]interface{}, error)
	Get(ctx context.Context, key string) (string, error)
	Expire(ctx context.Context, key string, duration time.Duration) error
	Close()
}

type Pair struct {
	Key   string
	Value interface{}
}
