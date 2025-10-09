package output

import (
	"fmt"
	"math"
	"text/tabwriter"

	"github.com/cpustejovsky/toggltrack/record"
)

func OutputStats(w *tabwriter.Writer, work float64, start, current, goal record.Record, timeLeft float64) {
	//Calculate minLeft
	initialMinutes := start.TotalMinutes()
	currentMinutes := current.TotalMinutes() + initialMinutes
	athMinutes := goal.TotalMinutes()
	goalPercentage := (currentMinutes / athMinutes) * 100
	goalHr := goal.TotalMinutes() / 60.0
	goalMin := math.Mod(goal.TotalMinutes(), 60.0)
	fmt.Fprintf(w, "%s\t%.0fh %.0fm\t\n", goal.Name(), math.Floor(goalHr), goalMin)

	curHr := math.Floor(currentMinutes / 60)
	curMin := math.Mod(currentMinutes, 60.0)
	fmt.Fprintf(w, "Today\t%.0fh %.0fm\t\n", math.Floor(current.TotalMinutes()/60), math.Mod(current.TotalMinutes(), 60.0))
	fmt.Fprintf(w, "Work Done\t%.0fh %.0fm\t%.1f%%\t\n", math.Floor(curHr), curMin, goalPercentage)

	if currentMinutes > athMinutes {
		t := currentMinutes - athMinutes
		fmt.Fprintf(w, "Extra Work\t%dhr %dm \t%.1f%%\t\n",
			int(t)/60,
			int(t)%60,
			goalPercentage-100)
	} else {
		//Calculate weekdays and weekend days
		gapMin := athMinutes - initialMinutes
		workDone := (currentMinutes - initialMinutes)

		var minLeft float64
		if workDone < work {
			minLeft = (work - workDone)
			fmt.Printf("Do %dhr %dm more work (%dmin)\n",
				int(minLeft)/60, int(minLeft)%60, int(minLeft))
			fmt.Printf("That's x<=%d pomodoros\n",
				int(math.Ceil(minLeft/25.0)))
		} else if workDone == work {
			fmt.Println("you've done all the work you needed to do today!")
		} else {
			minLeft = (workDone - work)
			fmt.Printf("you've done %dhr %dm of extra work!\n",
				int(minLeft)/60, int(minLeft)%60)
		}
		fmt.Fprintf(w, "Work Left\t%.0fh %.0fm\t%.1f%%\t\n", math.Floor((gapMin-workDone)/60.0), math.Mod((gapMin-workDone), 60.0), 100-goalPercentage)
		timeLeft -= (work - workDone)
		if timeLeft > 0 {
			fmt.Fprintf(w, "Time Left\t%.0fh %.0fm\n", math.Floor(timeLeft/60.0), math.Mod(timeLeft, 60.0))
		}
	}
	w.Flush()
}
