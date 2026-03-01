package calculator_test

import (
	"testing"
	"time"

	"github.com/cpustejovsky/track/calculator"
)

func TestDaysRemainingInMonth(t *testing.T) {
	expectWeekdays := 21.0
	expectWeekendDays := 10.0
	march := time.Date(2025, time.March, 1, 0, 0, 0, 0, time.UTC)
	gotWeekdays := calculator.WeekdaysRemaining(1, march)
	gotWeekendDays := calculator.WeekendDaysRemaining(1, march)
	if gotWeekendDays != expectWeekendDays {
		t.Logf("expected %v; got %v", expectWeekendDays, gotWeekendDays)
	}
	if gotWeekdays != expectWeekdays {
		t.Logf("expected %v; got %v", expectWeekdays, gotWeekdays)
	}
}

func TestCalculateWorkToday(t *testing.T) {
}
