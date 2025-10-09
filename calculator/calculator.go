package calculator

import (
	"math"
	"time"
)

type Calculator struct {
	Crunch          bool
	WeekendWorkTime float64
}

func Foobar() {

}

func New(crunch bool, weekendWorkTime float64) Calculator {
	return Calculator{
		Crunch:          crunch,
		WeekendWorkTime: weekendWorkTime,
	}
}

func daysInMonth(date time.Time) int {
	return time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func isWeekDay(d time.Weekday) bool {
	return d >= time.Monday && d <= time.Friday
}

func WeekdaysRemaining(start int, date time.Time) float64 {
	var weekDays float64
	end := daysInMonth(date)
	for i := start; i <= end; i++ {
		next := time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, time.UTC)
		if isWeekDay(next.Weekday()) {
			weekDays++
		}
	}
	return weekDays
}
func WeekendDaysRemaining(start int, date time.Time) float64 {
	var weekendDays float64
	end := daysInMonth(date)
	for i := start; i <= end; i++ {
		next := time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, time.UTC)
		if !isWeekDay(next.Weekday()) {
			weekendDays++
		}
	}
	return weekendDays
}

func (c Calculator) CalculateWorkToday(gap float64) float64 {
	now := time.Now()
	weekDays := WeekdaysRemaining(now.Day(), now)
	weekEndDays := WeekendDaysRemaining(now.Day(), now)
	if c.Crunch {
		return math.Ceil(gap / (weekEndDays + weekDays))
	}
	if isWeekDay(now.Weekday()) {
		return math.Ceil((gap - c.WeekendWorkTime*weekEndDays) / weekDays)
	} else {
		return c.WeekendWorkTime
	}
}

func (c Calculator) CalculateWorkWeekDay(gap float64) float64 {
	now := time.Now()
	weekDays := WeekdaysRemaining(now.Day()+1, now)
	weekEndDays := WeekendDaysRemaining(now.Day()+1, now)
	if c.Crunch {
		return math.Ceil(gap / (weekEndDays + weekDays))
	}
	weekendWork := c.WeekendWorkTime * weekEndDays
	return math.Ceil((gap - weekendWork) / weekDays)

}
