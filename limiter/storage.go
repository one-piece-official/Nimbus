package limiter

import (
	"context"
	"time"
)

type LimitStorage interface {
	Incr(context context.Context, key string, count int64) (int64, error)
	Expire(context context.Context, key string, ex time.Duration) error
	Get(context context.Context, key string) (int64, error)
}
