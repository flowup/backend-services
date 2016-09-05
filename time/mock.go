package time

import "time"

//Service is an interface for mock time
type Service interface {
	SetTime(time time.Time)
	Now() time.Time
	Sleep(duration time.Duration)
	Rollback(duration time.Duration)
}
