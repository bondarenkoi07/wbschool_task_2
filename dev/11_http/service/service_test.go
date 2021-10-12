package service

import (
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/calendar"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	var service = NewService(calendar.NewCalendar())
	var TestDates = map[bool][]string{
		true:  {"2000-1-12", "2000-09-01", "2021-31-12"},
		false: {"2000-32-02", "2000-1100-22322", "", "eafafafafaffaaf"},
	}

	for b, strings := range TestDates {
		for _, s := range strings {
			if b {
				_, err := service.parse("test", s)
				if err != nil {
					t.Errorf("valid date return not-nil err: %v", err)
				}
			} else {
				_, err := service.parse("test", s)
				if err == nil {
					t.Errorf("invalid date return nil err, %s", s)
				}
			}
		}
	}
}

func TestFilters(t *testing.T) {
	var service = NewService(calendar.NewCalendar())
	testData := []calendar.Event{
		{
			Date:        time.Now(),
			Description: "2021-12-10",
		},
		{
			Date:        time.Now().AddDate(0, 1, 0),
			Description: "2021-12-11",
		},
		{
			Date:        time.Now().AddDate(0, 0, 1),
			Description: "2021-13-10",
		},
	}
	for _, datum := range testData {
		err := service.Create(datum.Description, datum.Description)
		if err != nil {
			t.Error(err)
		}
	}

	_, err := service.FilterByDay("2021-12-10")
	if err != nil {
		t.Error(err)
	}

	_, err = service.FilterByMonth("2021-12-10")
	if err != nil {
		t.Error(err)
	}

	_, err = service.FilterByWeek("2021-12-10")
	if err != nil {
		t.Error(err)
	}
}
