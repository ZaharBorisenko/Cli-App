package handlers

import (
	"fmt"
	"time"
)

type TimeManager struct {
	todos *Todos
}

func NewTimeManager(todos *Todos) *TimeManager {
	return &TimeManager{todos: todos}
}

func (tm *TimeManager) SetDateCompleted(id int, date string) error {
	if err := (*tm.todos).ValidateId(id); err != nil {
		return err
	}
	result, err := tm.ValidateDate(date)
	if err != nil {
		return err
	}

	(*tm.todos)[id].Deadline = result
	return nil
}
func (tm *TimeManager) ValidateDate(date string) (*time.Time, error) {
	layout := "02/01/2006"
	currentTime := time.Now().Local()

	t, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		return nil, fmt.Errorf("invalid date format, use DD/MM/YYYY: %v", err)
	}

	if currentTime.After(t) && !isSameDay(currentTime, t) {
		return nil, fmt.Errorf("cannot set past date: %s", date)
	}

	return &t, nil
}

func isSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day()
}
