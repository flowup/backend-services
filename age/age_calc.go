package age

import "time"

// AgeService is a type for the interface Service
type AgeService struct {
}

// NewAgeService is a constructor creating new AgeService
func NewAgeService() *AgeService {
	return &AgeService{}
}

// CalculateAge is function calculating time period from the given event to now
func (s *AgeService) CalculateAge(event time.Time, now time.Time) time.Duration {

	duration := now.Sub(event)
	return duration
}
