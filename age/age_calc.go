package age

import "time"

// AgeService is a struct
type AgeService struct {
}

// NewAgeService is a constructor creating new AgeService
func NewAgeService() *AgeService {
	return &AgeService{}
}

// CalculateAge is function calculating years from the given event to today
func (s *AgeService) CalculateAge(event time.Time, now time.Time) time.Duration {

	/*
		ABANDONED: Precise calculation of years

		now := time.Now()
		years := now.Year() - event.Year()

		//	Checks if event is going to happen this year but next month or in next days in current month
		//	if so one year must be decremented
		if now.Month() < event.Month() {
			years--
		} else if now.Month() == event.Month() && now.Day() < event.Day() {
			years--
		}
	*/

	duration := now.Sub(event)
	return duration
}
