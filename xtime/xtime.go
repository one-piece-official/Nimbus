package xtime

import (
	"fmt"
	"time"
)

const (
	DefaultDate     = "2006-01-02"
	DefaultTime     = "15:04:05"
	DefaultDateTime = "2006-01-02 15:04:05"
)

var (
	halfAnHour      = 30
	cisolarixSunday = 7
)

type XTime time.Time

// NOTE: 为了兼容 Rails 的后台专门做的时间戳模式.
func CisolarixTimeStr(t time.Time) string {
	timeStr := fmt.Sprintf("%d-%d", weekday(t), t.Hour())
	if t.Minute() >= halfAnHour {
		timeStr += ".5"
	}

	return timeStr
}

// NOTE: cisolarix 的周日是星期 7.
func weekday(t time.Time) int {
	weekday := t.Weekday()

	if weekday == time.Sunday {
		return cisolarixSunday
	}

	return int(weekday)
}

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func TodayDateStr() string {
	return time.Now().Format(DefaultDate)
}

func TodayTimeStr() string {
	return time.Now().Format(DefaultTime)
}

func TodayDateTimeStr() string {
	return time.Now().Format(DefaultDateTime)
}

func TodayStart() int64 {
	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	return date.UnixNano() / int64(time.Millisecond)
}

func Time(days int) XTime {
	date := time.Now()

	return XTime(time.Date(date.Year(), date.Month(), date.Day()+days, 0, 0, 0, 0, date.Location()))
}

func (xtime XTime) DayStart() int64 {
	t := time.Time(xtime)

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixNano() / int64(time.Millisecond)
}

func (xtime XTime) DateStr() string {
	return time.Time(xtime).Format(DefaultDate)
}

func (xtime XTime) TimeStr() string {
	return time.Time(xtime).Format(DefaultTime)
}

func (xtime XTime) DateTimeStr() string {
	return time.Time(xtime).Format(DefaultDateTime)
}
