package repository_test

import (
	"context"
	"testing"

	"github.com/one-piece-official/Nimbus/repository"
	"github.com/stretchr/testify/assert"
)

func TestIncr(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mapKV := repository.NewMapKV(make(map[string]interface{}))

	err := mapKV.Incr(ctx, "1")
	assert.Nil(t, err)

	value, err := mapKV.Get(ctx, "1")
	assert.Equal(t, value, "1")
	assert.Nil(t, err)

	err = mapKV.Incr(ctx, "1")
	assert.Nil(t, err)

	value, err = mapKV.Get(ctx, "1")
	assert.Equal(t, value, "2")
	assert.Nil(t, err)

	_, err = mapKV.IncrByAndGet(ctx, "1", -2)
	assert.Nil(t, err)

	value, err = mapKV.Get(ctx, "1")
	assert.Equal(t, value, "0")
	assert.Nil(t, err)
}

func TestSetAndGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mapKV := repository.NewMapKV(make(map[string]interface{}))

	value, err := mapKV.Get(ctx, "1")
	assert.Equal(t, value, "")
	assert.NotNil(t, err)

	_ = mapKV.Set(ctx, "1", 1)

	value, err = mapKV.Get(ctx, "1")
	assert.Equal(t, value, "1")
	assert.Nil(t, err)

	_ = mapKV.Set(ctx, "1", "1")

	value, err = mapKV.Get(ctx, "1")
	assert.Equal(t, value, "1")
	assert.Nil(t, err)

	_, _ = mapKV.Del(ctx, "1")

	value, err = mapKV.Get(ctx, "1")
	assert.Equal(t, value, "")
	assert.NotNil(t, err)
}
