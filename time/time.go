package time

import "time"

// Time is mock of time
type Time struct {
	time time.Time
}

// NewMockTime will create new mock time by given time
func NewMockTime(time time.Time) *Time {
	return &Time{time: time}
}

// SetTime will set mock time by given time
func (t *Time) SetTime(time time.Time) {
	t.time = time
}

// Now will return mock time
func (t *Time) Now() time.Time {
	return t.time
}

// Sleep will add duration to time
func (t *Time) Sleep(duration time.Duration) {
	t.time = t.time.Add(duration)
}

// Rollback will go back in time by duration
func (t *Time) Rollback(duration time.Duration) {
	t.time = t.time.Add(-duration)
}
