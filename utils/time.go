package utils

import (
	"time"
)

func BeginOfDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, timePoint.Location())
}

func StartOfNextDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, timePoint.Location())
}
