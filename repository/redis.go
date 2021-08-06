package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	mSetPair = 2
)

type RedisKV struct {
	db *redis.Client
}

func NewRedisKv(db *redis.Client) *RedisKV {
	return &RedisKV{db: db}
}

func (r *RedisKV) Incr(ctx context.Context, key string) error {
	return r.db.Incr(ctx, key).Err()
}

func (r *RedisKV) IncrAndGet(ctx context.Context, key string) (int64, error) {
	return r.db.Incr(ctx, key).Result()
}

func (r *RedisKV) IncrByAndGet(ctx context.Context, key string, value int64) (int64, error) {
	return r.db.IncrBy(ctx, key, value).Result()
}

func (r *RedisKV) Expire(ctx context.Context, key string, duration time.Duration) error {
	return r.db.Expire(ctx, key, duration).Err()
}

func (r *RedisKV) Set(ctx context.Context, key string, value interface{}) error {
	return r.db.Set(ctx, key, value, 0).Err()
}

func (r *RedisKV) SetWithTTL(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.db.SetEX(ctx, key, value, expiration).Err()
}

func (r *RedisKV) Del(ctx context.Context, key string) (bool, error) {
	cmd := r.db.Del(ctx, key)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}

	return cmd.Val() == 1, nil
}

func (r *RedisKV) MSet(ctx context.Context, pair ...Pair) error {
	values := make([]interface{}, len(pair)*mSetPair)
	for i := range pair {
		values[2*i] = pair[i].Key
		values[2*i+1] = pair[i].Value
	}

	cmd := r.db.MSet(ctx, values)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *RedisKV) MSetWithTTL(ctx context.Context, expiration time.Duration, pair ...Pair) error {
	_, err := r.db.Pipelined(ctx, func(p redis.Pipeliner) error {
		for i := range pair {
			if _, err := p.Set(ctx, pair[i].Key, pair[i].Value, expiration).Result(); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to run MSetWithTTL: %w", err)
	}

	return nil
}

func (r *RedisKV) HIncr(ctx context.Context, key, field string, incr int64) error {
	return r.db.HIncrBy(ctx, key, field, incr).Err()
}

func (r *RedisKV) Exists(ctx context.Context, key string) (bool, error) {
	ret, err := r.db.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to run Exists: %w", err)
	}

	return ret == 1, nil
}

func (r *RedisKV) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return r.db.MGet(ctx, keys...).Result()
}

func (r *RedisKV) Get(ctx context.Context, key string) (string, error) {
	value, err := r.db.Get(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err = ErrorKVNil
	}

	return value, err
}

func (r *RedisKV) Close() {
	_ = r.db.Close()
}
