package age

import "time"

// Service is the interface of ageService implementation
type Service interface {
	CalculateAge(event time.Time, now time.Time) time.Duration
}
