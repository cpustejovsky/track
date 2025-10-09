package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	"github.com/cpustejovsky/toggltrack/calculator"
	"github.com/cpustejovsky/toggltrack/flags"
	"github.com/cpustejovsky/toggltrack/record"
)

var (
	compareYear         = flags.Int("compareYear", "COMPARE_YEAR", 1971, "The month you are comparing to (1-12)")
	compareMonth        = flags.Int("compareMonth", "COMPARE_MONTH", 1, "The month you are comparing to (1-12)")
	compareMonth_hour   = flags.Int("compareMonth_hour", "COMPARE_MONTH_HOUR", 1, "Total hours for your most worked month")
	compareMonth_minute = flags.Int("compareMonth_minute", "COMPARE_MONTH_MINUTE", 0, "Total minutes mod 60 for your most worked month")
	crunch              = flags.Bool("crunch", "CRUNCH", false, "Whether you want to work without weekend breaks")
	ideal               = flags.Float64("ideal", "IDEAL", 1.05, "The percentage you want to grow")
	weekendWork         = flags.Float64("weekendWork", "WEEKEND_WORK", 0.0, "Total minutes you want to work on Saturdays and Sundays")
)

func main() {
	flag.Parse()
	date := time.Date(*compareYear, time.Month(*compareMonth), 1, 0, 0, 0, 0, time.UTC)
	compareMonth := record.New(fmt.Sprintf("%s %d", date.Month(), date.Year()), *compareMonth_hour, *compareMonth_minute)
	calc := calculator.New(*crunch, *weekendWork, *ideal)
	i := calc.CalculateIdeal(compareMonth.TotalMinutes(), date)
	fmt.Println("Ideal", int(math.Round(i/60.0)))
}
