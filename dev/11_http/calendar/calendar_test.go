package calendar

import (
	"testing"
	"time"
)

func TestCalendar_Filter(t *testing.T) {
	var calendar = NewCalendar()
	testData := []Event{
		{
			Date:        time.Now(),
			Description: "",
		},
		{
			Date:        time.Now().AddDate(0, 1, 0),
			Description: "",
		},
		{
			Date:        time.Now().AddDate(0, 0, 1),
			Description: "",
		},
	}
	for _, datum := range testData {
		calendar.Create(datum)
	}

	if len(calendar.ReadAll()) != len(testData) {
		t.Errorf("wrong calendar: %v", calendar.ReadAll())
	}

	filtered := calendar.FilterByDay(time.Now())
	if len(filtered) != 1 {
		t.Errorf("wrong filter by day %v", filtered)
	}

	filtered = calendar.FilterByMonth(time.Now().AddDate(0, 0, -time.Now().Day()+1))
	if len(filtered) != 2 {
		t.Errorf("wrong filter by month %v", filtered)
	}

	filtered = calendar.FilterByWeek(time.Now().AddDate(0, 0, int(-time.Now().Weekday())))
	if len(filtered) != 2 {
		t.Errorf("wrong filter by week %v", filtered)
	}
}
