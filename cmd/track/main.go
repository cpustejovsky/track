package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/cpustejovsky/toggltrack/calculator"
	"github.com/cpustejovsky/toggltrack/flags"
	"github.com/cpustejovsky/toggltrack/output"
	"github.com/cpustejovsky/toggltrack/record"
)

var (
	compareYear         = flags.Int("compareYear", "COMPARE_YEAR", 1971, "The month you are comparing to (1-12)")
	compareMonth        = flags.Int("compareMonth", "COMPARE_MONTH", 1, "The month you are comparing to (1-12)")
	compareMonth_hour   = flags.Int("compareMonth_hour", "COMPARE_MONTH_HOUR", 1, "Total hours for your most filled month")
	compareMonth_minute = flags.Int("compareMonth_minute", "COMPARE_MONTH_MINUTE", 0, "Total minutes mod 60 for your most filled month")
	crunch              = flags.Bool("crunch", "CRUNCH", false, "Whether you want to fill your time without weekend breaks")
	showCompareMonth    = flags.Bool("showCompareMonth", "SHOW_COMPARE_MONTH", true, "Whether you want to show the month you're comparing to")
	ideal               = flags.Int("ideal", "IDEAL", 2, "The ideal amount of hours you want to fill today")
	showIdeal           = flags.Bool("showIdeal", "SHOW_IDEAL", true, "Whether you want to show your ideal goal")
	weekendWork         = flags.Float64("weekendWork", "WEEKEND_WORK", 0.0, "Total minutes you want to fill on Saturdays and Sundays")
	endOfDayHr          = flags.Int("eod_hr", "EOD_HR", 21, "Hour your day ends (0 to 24 hours)")
	endOfDayMin         = flags.Int("eod_min", "EOD_MIN", 0, "Minute your day ends (0 to 59 hours)")
)

func argsToInts(args ...string) []int {
	nums := make([]int, len(args))
	for i, arg := range args {
		x, err := strconv.Atoi(arg)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		nums[i] = x
	}
	return nums
}

func main() {
	flag.Parse()
	args := flag.Args()
	date := time.Date(*compareYear, time.Month(*compareMonth), 1, 0, 0, 0, 0, time.UTC)
	compareMonth := record.New(fmt.Sprintf("%s %d", date.Month(), date.Year()), *compareMonth_hour, *compareMonth_minute)
	now := time.Now()
	calc := calculator.New(*crunch, *weekendWork)
	Ideal := record.New("Ideal", *ideal, 00)
	if len(args) < 4 {
		log.Println("please provide hours and minutes")
		os.Exit(1)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	//Parse arguments
	nums := argsToInts(args...)
	nowName := fmt.Sprintf("%s %d", now.Month(), now.Year())
	start := record.New(nowName, (nums[0]), (nums[1]))
	current := record.New(nowName, (nums[2]), (nums[3]))
	// TODO: determine whether to keep or set conditionally
	// fmt.Printf("Current time is %d:%02d\n", now.Hour(), now.Minute())
	compare := compareMonth
	var work float64
	if *showIdeal {
		work = Ideal.TotalMinutes()
	} else {
		work = calc.CalculateWorkToday(compareMonth.TotalMinutes() - start.TotalMinutes())
	}
	eod := time.Date(now.Year(), now.Month(), now.Day(), *endOfDayHr, *endOfDayMin, 0, 0, now.Location())
	timeLeft := eod.Sub(now).Minutes()
	output.OutputStats(w, work, start, current, compare, timeLeft)
}
