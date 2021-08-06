package limiter

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type Limiter interface {
	Incr() (int64, error)
	Available() bool
}

type baseLimiter struct {
	ctx       context.Context
	storage   LimitStorage
	key       string
	duration  time.Duration
	threshold int64
}

func NewFrequencyLimiter(ctx context.Context, storage LimitStorage, key string,
	duration time.Duration, threshold int64) Limiter {
	return &frequencyLimiter{baseLimiter{
		ctx:       ctx,
		storage:   storage,
		key:       key,
		duration:  duration,
		threshold: threshold,
	}}
}

func NewThresholdLimiter(ctx context.Context, storage LimitStorage, key string,
	duration time.Duration, threshold int64) Limiter {
	return &thresholdLimiter{baseLimiter{
		ctx:       ctx,
		storage:   storage,
		key:       key,
		duration:  duration,
		threshold: threshold,
	}}
}

type frequencyLimiter struct {
	baseLimiter
}

type thresholdLimiter struct {
	baseLimiter
}

func (f baseLimiter) Incr() (int64, error) {
	val, err := f.storage.Incr(f.ctx, f.key, 1)
	if err != nil {
		return 0, errors.Wrap(err, "storage incr err")
	}

	if val <= 1 && f.duration > 0 {
		err = f.storage.Expire(f.ctx, f.key, f.duration)
		if err != nil {
			return val, errors.Wrap(err, "storage expire err")
		}
	}

	return val, nil
}

func (f frequencyLimiter) Available() bool {
	val, err := f.storage.Get(f.ctx, f.key)
	if err != nil {
		return false
	}

	return val == 1 || val%f.threshold == 0
}

func (f thresholdLimiter) Available() bool {
	val, err := f.storage.Get(f.ctx, f.key)
	if err != nil {
		return false
	}

	return val < f.threshold
}
