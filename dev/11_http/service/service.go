package service

import (
	"errors"
	"fmt"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/calendar"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/serializers"
	"strconv"
	"strings"
	"time"
)

const ParseDateErrorFormat = "InputError: invalid date format: got %s, must be (year)-(day)-(month)"

type Service struct {
	calendar *calendar.Calendar
}

func NewService(calendar *calendar.Calendar) *Service {
	return &Service{calendar: calendar}
}

func (s *Service) parse(description, date string) (calendar.Event, error) {
	var (
		event               calendar.Event
		day, year, RawMonth int
		err                 error
	)

	dateParts := strings.Split(date, "-")

	if len(dateParts) != 3 {
		return calendar.Event{}, errors.New(fmt.Sprintf(ParseDateErrorFormat, date))
	}

	year, err = strconv.Atoi(dateParts[0])
	if err != nil || year <= 0 {
		return calendar.Event{}, errors.New(fmt.Sprintf(ParseDateErrorFormat, date))
	}

	RawMonth, err = strconv.Atoi(dateParts[2])
	if err != nil || RawMonth <= 0 || RawMonth >= 13 {
		return calendar.Event{}, errors.New(fmt.Sprintf(ParseDateErrorFormat, date))
	}

	day, err = strconv.Atoi(dateParts[1])
	if err != nil || day <= 0 || day > 31 {
		return calendar.Event{}, errors.New(fmt.Sprintf(ParseDateErrorFormat, date))
	}

	month := time.Month(RawMonth)

	firstOfMonth := time.Date(
		year, month, 1, 0, 0, 0, 0, time.Local,
	)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	if lastOfMonth.Day() < day {
		return calendar.Event{}, errors.New(fmt.Sprintf(ParseDateErrorFormat, date))
	}

	event.Date = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	event.Description = description
	return event, nil
}

func (s *Service) Update(id uint, description, date string) error {
	event, err := s.parse(description, date)

	if err != nil {
		return err
	}

	err = (*s).calendar.Update(id, event)
	return err
}

func (s *Service) Create(description, date string) error {
	event, err := s.parse(description, date)

	if err != nil {
		return err
	}

	(*s).calendar.Create(event)

	return nil
}

func (s *Service) Delete(id uint) error {
	err := (*s).calendar.Delete(id)
	return err
}

func (s *Service) Read(id uint) (string, error) {
	read, err := (*s).calendar.Read(id)
	return serializers.SerializeSingleEvent(read), err
}

func (s *Service) ReadAll() string {
	return serializers.SerializeMap((*s).calendar.ReadAll())
}

func (s *Service) FilterByDay(date string) (string, error) {
	parsed, err := s.parse("", date)
	if err != nil {
		return "", err
	}
	dateParsed := parsed.Date
	return serializers.SerializeSlice((*s).calendar.FilterByDay(dateParsed)), nil
}

func (s *Service) FilterByWeek(WeekFirstDayDate string) (string, error) {
	parsed, err := s.parse("", WeekFirstDayDate)
	if err != nil {
		return "", err
	}

	dateParsed := parsed.Date.AddDate(0, 0, int(-parsed.Date.Weekday()))
	return serializers.SerializeSlice((*s).calendar.FilterByWeek(dateParsed)), nil
}

func (s *Service) FilterByMonth(monthDate string) (string, error) {
	parsed, err := s.parse("", monthDate)
	if err != nil {
		return "", err
	}
	dateParsed := parsed.Date

	firstOfMonth := time.Date(
		dateParsed.Year(), dateParsed.Month(), 1, 0, 0, 0, 0, dateParsed.Location(),
	)

	return serializers.SerializeSlice((*s).calendar.FilterByMonth(firstOfMonth)), nil
}
