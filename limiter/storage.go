package limiter

import (
	"context"
	"strconv"
	"time"

	"github.com/one-piece-official/Nimbus/repository"
	"github.com/pkg/errors"
)

type LimitStorage interface {
	Incr(context context.Context, key string, count int64) (int64, error)
	Expire(context context.Context, key string, ex time.Duration) error
	Get(context context.Context, key string) (int64, error)
}

type redisLimitStorage struct {
	kvDB repository.KVIface
}

func (r redisLimitStorage) Expire(context context.Context, key string, ex time.Duration) error {
	return r.kvDB.Expire(context, key, ex)
}

func (r redisLimitStorage) Incr(context context.Context, key string, count int64) (int64, error) {
	return r.kvDB.IncrAndGet(context, key)
}

// TODO 是否需要支持 mget.
func (r redisLimitStorage) Get(context context.Context, key string) (int64, error) {
	value, err := r.kvDB.Get(context, key)
	if err != nil {
		return 0, errors.Wrap(err, "cannot get from redis")
	}

	count, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.Wrap(err, "value is not int")
	}

	return int64(count), nil
}

func NewRedisLimitStorage(kvDB repository.KVIface) LimitStorage {
	return &redisLimitStorage{
		kvDB: kvDB,
	}
}
