package limiter_test

import (
	"context"
	"testing"
	"time"

	"github.com/one-piece-official/Nimbus/limiter"

	"github.com/one-piece-official/Nimbus/repository"
	"github.com/stretchr/testify/assert"
)

func TestLimiterAvailable(t *testing.T) {
	t.Parallel()

	// ctx := context.Background()

	limiterKV := repository.NewMapKV(make(map[string]interface{}))

	limiterObj := limiter.NewThresholdLimiter(context.Background(), limiter.NewRedisLimitStorage(limiterKV), "", time.Hour, 2)

	checkLimiter := limiterObj.Available()
	assert.Equal(t, true, checkLimiter)

	_, _ = limiterObj.Incr()
	checkLimiter = limiterObj.Available()
	assert.Equal(t, true, checkLimiter)

	_, _ = limiterObj.Incr()
	checkLimiter = limiterObj.Available()
	assert.Equal(t, false, checkLimiter)
}
