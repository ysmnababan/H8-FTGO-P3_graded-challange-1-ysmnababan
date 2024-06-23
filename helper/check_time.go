package helper

import (
	"time"
)

func IsTimeValid(startTime, endTime string) bool {
	// convert to time.Time object
	layout := "2006-01-02 15:04:05"

	// Parse the string according to the layout
	start, err := time.Parse(layout, startTime)
	if err != nil {
		Logging(nil).Error(err)
		return false
	}

	end, err := time.Parse(layout, endTime)
	if err != nil {
		Logging(nil).Error(err)
		return false
	}
	// fmt.pr
	// both time must be in the same day
	if !isSameDay(start, end) {
		return false
	}
	// rental min time : 1 hour
	if start.Hour() >= end.Hour() {
		return false
	}
	return true
}

func isSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}
