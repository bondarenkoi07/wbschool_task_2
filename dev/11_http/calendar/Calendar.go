package calendar

import (
	"errors"
	"sync"
	"time"
)

type Calendar struct {
	sync.RWMutex
	events map[uint]Event
	id     uint
}

func NewCalendar() *Calendar {
	return &Calendar{events: make(map[uint]Event)}
}

func (c *Calendar) Delete(id uint) error {
	var err error
	(*c).Lock()
	_, ok := c.events[id]
	if ok {
		delete((*c).events, id)
	} else {
		err = errors.New("logic:there is nothing to delete")
	}
	(*c).Unlock()

	return err
}

func (c *Calendar) Read(id uint) (Event, error) {
	(*c).RLock()
	defer (*c).RUnlock()
	val, ok := c.events[id]
	if ok {
		return val, nil
	} else {
		return Event{}, errors.New("logic:not found")
	}
}

func (c *Calendar) ReadAll() map[uint]Event {
	(*c).RLock()
	output := c.events
	(*c).RUnlock()
	return output
}

func (c *Calendar) Create(event Event) {
	(*c).Lock()
	(*c).events[c.id] = event
	(*c).id++
	(*c).Unlock()
}

func (c *Calendar) Update(id uint, event Event) error {
	(*c).Lock()
	defer (*c).Unlock()
	_, ok := c.events[id]
	if ok {
		(*c).events[id] = event
	} else {
		return errors.New("logic: nothing to update")
	}
	(*c).id++
	return nil
}

func (c *Calendar) FilterByDay(date time.Time) []Event {
	NextDay := date.Local().AddDate(0, 0, 1)
	output := make([]Event, 0, len(c.events))
	(*c).RLock()
	for _, day := range c.events {
		if NextDay.After(day.Date) && date.Before(NextDay) {
			output = append(output, day)
		}
	}
	(*c).RUnlock()
	return output[:]
}

func (c *Calendar) FilterByWeek(date time.Time) []Event {
	NextWeekFirstDay := date.Local().AddDate(0, 0, 6)
	output := make([]Event, 0, len(c.events))
	(*c).RLock()
	for _, day := range c.events {
		if NextWeekFirstDay.After(day.Date) && date.Before(day.Date) {
			output = append(output, day)
		}
	}
	(*c).RUnlock()
	return output[:]
}

func (c *Calendar) FilterByMonth(firstOfMonth time.Time) []Event {
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	output := make([]Event, 0, len(c.events))
	(*c).RLock()
	for _, day := range c.events {
		if lastOfMonth.After(day.Date) && firstOfMonth.Before(day.Date) {
			output = append(output, day)
		}
	}
	(*c).RUnlock()
	return output[:]
}

type Event struct {
	Date        time.Time
	Description string
}
