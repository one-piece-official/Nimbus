package xtime_test

import (
	"testing"
	"time"

	"github.com/one-piece-official/Nimbus/xtime"
	"github.com/stretchr/testify/assert"
)

const TimeLayout = "2006-01-02 15:04:05"

func TestCisolarixTimeStr(t *testing.T) {
	data := [3][2]string{
		{"2020-05-10 19:30:00", "7-19.5"},
		{"2020-05-10 19:20:00", "7-19"},
		{"2020-05-05 16:20:00", "2-16"},
	}

	for _, val := range data {
		str, _ := time.Parse(TimeLayout, val[0])
		t.Log(xtime.CisolarixTimeStr(str))
		assert.Equal(t, xtime.CisolarixTimeStr(str), val[1])
	}
}

func TestDayStart(t *testing.T) {
	t.Log(xtime.Time(1).DayStart())
}
