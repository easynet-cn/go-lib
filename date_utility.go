package golib

import (
	"math"
	"time"
)

func FirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)

	return ZeroTime(d)
}

func LastDateOfMonth(d time.Time) time.Time {
	return FirstDateOfMonth(d).AddDate(0, 1, -1)
}

func ZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func FirstDateOfWeek(d time.Time) time.Time {
	offset := int(time.Monday - d.Weekday())

	if offset > 0 {
		offset = -6
	}

	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()).AddDate(0, 0, offset)
}

func LastDateOfWeek(d time.Time) time.Time {
	offset := int(7 - d.Weekday())

	if offset > 6 {
		offset = 0
	}

	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()).AddDate(0, 0, offset)
}

func WeekOfYear(d time.Time) int {
	return int(math.Ceil(LastDateOfWeek(d).Sub(LastDateOfWeek(time.Date(d.Year(), 1, 1, 0, 0, 0, 0, d.Location()))).Hours()/(24*7))) + 1
}

func AddDate(t time.Time, year, month, day int) time.Time {
	targetDate := t.AddDate(year, month, -t.Day()+1)
	targetDay := targetDate.AddDate(0, 1, -1).Day()

	if targetDay > t.Day() {
		targetDay = t.Day()
	}

	targetDate = targetDate.AddDate(0, 0, targetDay-1+day)

	return targetDate
}
