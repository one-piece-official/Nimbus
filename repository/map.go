package repository

import (
	"context"
	"fmt"
	"time"
)

var (
	errorFieldValueNotInt = fmt.Errorf("field not a int value")
	errorValueNotInt      = fmt.Errorf("not a int value")
	errorValueNotHash     = fmt.Errorf("not a hash value")
)

type MapKV struct {
	db map[string]interface{}
}

func (r *MapKV) TTL(ctx context.Context, key string) (time.Duration, error) {
	if value, ok := r.db[key]; ok {
		data, dataOK := value.(map[string]interface{})
		if !dataOK {
			return time.Duration(-1), ErrorKVNoExpire
		}

		if expire, expireOK := data["expire"].(time.Duration); expireOK {
			return expire, nil
		} else {
			return time.Duration(-1), ErrorKVNoExpire
		}
	} else {
		return time.Duration(-2), ErrorKVNil
	}
}

func NewMapKV(db map[string]interface{}) *MapKV {
	return &MapKV{db: db}
}

func (r *MapKV) Incr(ctx context.Context, key string) error {
	currentValue := r.db[key]
	if currentValue == nil {
		currentValue = int64(0)
	}

	intValue, ok := currentValue.(int64)
	if !ok {
		return fmt.Errorf("%w %v", errorValueNotInt, currentValue)
	}

	r.db[key] = intValue + 1

	return nil
}

func (r *MapKV) IncrAndGet(ctx context.Context, key string) (int64, error) {
	if err := r.Incr(ctx, key); err != nil {
		return 0, err
	}

	value, ok := r.db[key].(int64)
	if !ok {
		return 0, fmt.Errorf("%w %v", errorValueNotInt, r.db[key])
	}

	return value, nil
}

func (r *MapKV) IncrByAndGet(ctx context.Context, key string, value int64) (int64, error) {
	currentValue := r.db[key]
	if currentValue == nil {
		currentValue = int64(0)
	}

	intValue, ok := currentValue.(int64)
	if !ok {
		return 0, fmt.Errorf("%w %v", errorValueNotInt, currentValue)
	}

	r.db[key] = intValue + value

	newValue, ok := r.db[key].(int64)
	if !ok {
		return 0, fmt.Errorf("%w %v", errorValueNotInt, r.db[key])
	}

	return newValue, nil
}

func (r *MapKV) DecrByAndGet(ctx context.Context, key string, value int64) (int64, error) {
	currentValue := r.db[key]
	if currentValue == nil {
		currentValue = int64(0)
	}

	intValue, ok := currentValue.(int64)
	if !ok {
		return 0, fmt.Errorf("%w %v", errorValueNotInt, currentValue)
	}

	r.db[key] = intValue - value

	newValue, ok := r.db[key].(int64)
	if !ok {
		return 0, fmt.Errorf("%w %v", errorValueNotInt, r.db[key])
	}

	return newValue, nil
}

func (r *MapKV) Expire(ctx context.Context, key string, duration time.Duration) error {
	return nil
}

func (r *MapKV) Set(ctx context.Context, key string, value interface{}) error {
	r.db[key] = value

	return nil
}

func (r *MapKV) SetWithTTL(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	r.db[key] = value

	return nil
}

func (r *MapKV) Del(ctx context.Context, key string) (bool, error) {
	deleted := true
	if r.db[key] == nil {
		deleted = false
	}

	delete(r.db, key)

	return deleted, nil
}

func (r *MapKV) MSet(ctx context.Context, pair ...Pair) error {
	for _, kv := range pair {
		r.db[kv.Key] = kv.Value
	}

	return nil
}

func (r *MapKV) MSetWithTTL(ctx context.Context, expiration time.Duration, pair ...Pair) error {
	return r.MSet(ctx, pair...)
}

func (r *MapKV) HIncr(ctx context.Context, key, field string, incr int64) error {
	mapValue, ok := r.db[key].(map[string]interface{})
	if !ok {
		return fmt.Errorf("%w %v", errorValueNotHash, r.db[key])
	}

	intValue, ok := mapValue[field].(int64)
	if !ok {
		return fmt.Errorf("%w %v", errorFieldValueNotInt, mapValue[field])
	}

	mapValue[field] = intValue + 1
	r.db[key] = mapValue

	return nil
}

func (r *MapKV) HIncrAndGet(ctx context.Context, key, field string, incr int64) (int64, error) {
	mapValue, ok := r.db[key].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("%w %v", errorValueNotHash, r.db[key])
	}

	intValue, ok := mapValue[field].(int64)
	if !ok {
		return 0, fmt.Errorf("%w %v", errorFieldValueNotInt, mapValue[field])
	}

	mapValue[field] = intValue + 1
	r.db[key] = mapValue

	return intValue + 1, nil
}

func (r *MapKV) Exists(ctx context.Context, key string) (bool, error) {
	return r.db[key] != nil, nil
}

func (r *MapKV) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	values := make([]interface{}, len(keys))

	for i, key := range keys {
		var value string
		if r.db[key] != nil {
			if valueMap, ok := r.db[key].(map[string]interface{}); ok {
				value = fmt.Sprintf("%v", valueMap["value"])
			} else {
				value = fmt.Sprintf("%v", r.db[key])
			}
		}

		values[i] = value
	}

	return values, nil
}

func (r *MapKV) Get(ctx context.Context, key string) (string, error) {
	var err error

	var value string
	if r.db[key] != nil {
		value = fmt.Sprintf("%v", r.db[key])
	} else {
		err = ErrorKVNil
	}

	return value, err
}

func (r *MapKV) HGet(ctx context.Context, key, field string) (string, error) {
	var err error
	mapValue, ok := r.db[key].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("%w %v", errorValueNotHash, r.db[key])
	}

	var value string
	if mapValue[field] != nil {
		value = fmt.Sprintf("%v", mapValue[field])
	} else {
		err = ErrorKVNil
	}

	return value, err
}

func (r *MapKV) Close() {
}
