package limiter

import (
	"context"
	"fmt"
	"time"
)

const (
	LimitTypeThreshold = "threshold"
	LimitTypeFrequency = "frequency"

	LimitTimeTypeToday     = "today"
	LimitTimeTypeHour      = "hour"
	LimitTimeTypeHourToday = "hourToday"
	LimitTimeTypeForever   = "forever"

	hoursPerDay = 24
)

func getTimeKeyAndDuration(timeType string, limitTimeLength int) (timeKey string, duration time.Duration) {
	switch timeType {
	case LimitTimeTypeHour:
		return timeType, time.Hour * time.Duration(limitTimeLength)
	case LimitTimeTypeHourToday:
		return timeType + time.Now().Format("2006-01-02"), time.Hour * time.Duration(limitTimeLength)
	case LimitTimeTypeForever:
		return timeType, time.Hour * 366 * hoursPerDay
	default:
		return time.Now().Format("2006-01-02"), time.Hour * hoursPerDay
	}
}

func Available(ctx context.Context, storage LimitStorage, key, limitType, limitTimeType string, limitTimeLength int, limitAmount int64, needIncr bool) bool {
	limiter := GetLimiter(ctx, storage, key, limitType, limitTimeType, limitTimeLength, limitAmount)

	if needIncr {
		_, _ = limiter.Incr()

		return true
	}

	return limiter.Available()
}

func GetLimiter(ctx context.Context, storage LimitStorage, key, limitType, limitTimeType string, limitTimeLength int, limitAmount int64) Limiter {

	timeKey, duration := getTimeKeyAndDuration(limitTimeType, limitTimeLength)
	key = fmt.Sprintf("%s_%s", key, timeKey)

	switch limitType {
	case LimitTypeFrequency:
		return NewFrequencyLimiter(ctx, storage, key, duration, limitAmount)
	default:
		return NewThresholdLimiter(ctx, storage, key, duration, limitAmount)
	}
}
