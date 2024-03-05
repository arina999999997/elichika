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

func StartOfNextWeek(timePoint time.Time) time.Time {
	day := int(timePoint.Weekday()) // 0 for sunday, 1 for monday ...
	addedDay := 8 - day
	if addedDay == 8 {
		addedDay = 1
	}
	return BeginOfDay(timePoint.AddDate(0, 0, addedDay))
}
