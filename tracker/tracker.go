package tracker

import (
	"errors"
	"time"
)

// Errors
var (
	ErrorStart = errors.New("Trying to start event that was not stopped before.")
	ErrorStop  = errors.New("Trying to stop event that was not started before")
	ErrorNil   = errors.New("Type Event is nil")
)

// Event is struct that holds info about event
type Event struct {
	id      int64
	periods []Interval
}

// Tracker is a type for TrackerService
type Tracker struct {
	ID int64
}

// Interval is a type for TrackerService
type Interval struct {
	start time.Time
	end   time.Time
}

// NewTracker is a constructor creating new Tracker
func NewTracker() *Tracker {
	return &Tracker{ID: 1}
}

// StartNew is creating new Event. It is addding first interval to the array of periods of intervals.
func (s *Tracker) StartNew() *Event {

	interval := Interval{start: time.Now()}
	event := Event{id: s.ID, periods: []Interval{interval}}
	s.ID++
	return &event
}

// Start is taking existing event and adds a new interval to it.
func (s *Tracker) Start(e *Event) error {

	if e == nil {
		return ErrorNil
	}

	if e.periods[len(e.periods)-1].end.IsZero() == true {
		return ErrorStart
	}
	interval := Interval{start: time.Now()}
	e.periods = append(e.periods, interval)
	return nil
}

// Stop is stopping events
func (s *Tracker) Stop(e *Event) error {

	if e == nil {
		return ErrorNil
	}
	// If
	if e.periods[len(e.periods)-1].start.IsZero() == true {
		return ErrorStop
	}

	lastInterval := e.periods[len(e.periods)-1]
	e.periods = e.periods[:len(e.periods)-1]
	lastInterval.end = time.Now()
	e.periods = append(e.periods, lastInterval)

	return nil

}

// Total is summing all the events
func (s *Tracker) Total(e *Event) int64 {
	var duration int64
	for _, interval := range e.periods {
		if interval.end.IsZero() == true {
			duration += int64(time.Now().Sub(interval.start))
		} else {
			duration += int64(interval.end.Sub(interval.start))
		}
	}
	return duration

}
