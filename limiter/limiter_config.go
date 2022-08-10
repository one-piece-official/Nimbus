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
	LimitTimeTypeMinute    = "minute"
	LimitTimeTypeHourToday = "hourToday"
	LimitTimeTypeForever   = "forever"

	hoursPerDay = 24
)

func getTimeKeyAndDuration(timeType string, limitTimeLength int, hourDelta int64) (timeKey string, duration time.Duration) {
	now := time.Now().Add(time.Hour * time.Duration(hourDelta))

	switch timeType {
	case LimitTimeTypeHour:
		return timeType, time.Hour * time.Duration(limitTimeLength)
	case LimitTimeTypeMinute:
		return timeType, time.Minute * time.Duration(limitTimeLength)
	case LimitTimeTypeHourToday:
		return timeType + now.Format("2006-01-02"), time.Hour * time.Duration(limitTimeLength)
	case LimitTimeTypeForever:
		return timeType, time.Hour * 366 * hoursPerDay
	default:
		remainHours := hoursPerDay - now.Hour()

		return now.Format("2006-01-02"), time.Hour * time.Duration(remainHours)
	}
}

func Available(ctx context.Context, storage LimitStorage, key, limitType, limitTimeType string, limitTimeLength int, limitAmount, hourDelta int64, needIncr bool) bool {
	limiter := GetLimiter(ctx, storage, key, limitType, limitTimeType, limitTimeLength, limitAmount, hourDelta)

	if needIncr {
		_, _ = limiter.Incr()

		return true
	}

	return limiter.Available()
}

func GetLimiter(ctx context.Context, storage LimitStorage, key, limitType, limitTimeType string, limitTimeLength int, limitAmount int64, hourDelta int64) Limiter {
	timeKey, duration := getTimeKeyAndDuration(limitTimeType, limitTimeLength, hourDelta)
	key = fmt.Sprintf("%s_%s", key, timeKey)

	switch limitType {
	case LimitTypeFrequency:
		return NewFrequencyLimiter(ctx, storage, key, duration, limitAmount)
	default:
		return NewThresholdLimiter(ctx, storage, key, duration, limitAmount)
	}
}
