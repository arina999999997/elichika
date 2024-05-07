package utils

import (
	"time"
)

func BeginOfDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, timePoint.Location())
}

func BeginOfNextDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, timePoint.Location())
}

// return the next half day time point AFTER the current time point
// so when the time is [00:00, 12:00), return noon, otherwise return start of next day
func BeginOfNextHalfDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	noon := time.Date(year, month, day, 12, 0, 0, 0, timePoint.Location())
	if noon.After(timePoint) {
		return noon
	} else {
		return time.Date(year, month, day+1, 0, 0, 0, 0, timePoint.Location())
	}
}

// return the current half day time point
// so when the time is [00:00, 12:00), return start of day, otherwise return mid day
func BeginOfCurrentHalfDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	noon := time.Date(year, month, day, 12, 0, 0, 0, timePoint.Location())
	if noon.After(timePoint) {
		return time.Date(year, month, day, 0, 0, 0, 0, timePoint.Location())
	} else {
		return noon
	}
}

// return the next mid day AFTER the current time point
// so if the time is [00:00, 12:00), return the noon of today
// otherwise the turn the nood of next day
func NextMidDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	noon := time.Date(year, month, day, 12, 0, 0, 0, timePoint.Location())
	if noon.After(timePoint) {
		return noon
	} else {
		return time.Date(year, month, day+1, 12, 0, 0, 0, timePoint.Location())
	}
}

// return the next mid day AFTER the current time point
// so if the time is [00:00, 12:00), return the noon of today
// otherwise the turn the nood of next day
func CurrentMidDay(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	noon := time.Date(year, month, day, 12, 0, 0, 0, timePoint.Location())
	if noon.After(timePoint) {
		return time.Date(year, month, day-1, 12, 0, 0, 0, timePoint.Location())
	} else {
		return noon
	}
}

func StartOfNextWeek(timePoint time.Time) time.Time {
	day := int(timePoint.Weekday()) // 0 for sunday, 1 for monday ...
	addedDay := 8 - day
	if addedDay == 8 {
		addedDay = 1
	}
	return BeginOfDay(timePoint.AddDate(0, 0, addedDay))
}
