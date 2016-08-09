package age

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var testingDates = [3]time.Time{
	time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2013, 8, 2, 0, 0, 0, 0, time.UTC),
}

type AgeServiceSuite struct {
	suite.Suite
}

func (s *AgeServiceSuite) TestNewAgeService() {
	service := NewAgeService()
	assert.NotEqual(s.T(), nil, service)
}

func (s *AgeServiceSuite) TestCalculateAge() {
	service := NewAgeService()
	timeNow := time.Now()

	for _, testDate := range testingDates {
		assert.Equal(s.T(), timeNow.Sub(testDate), service.CalculateAge(testDate, timeNow))
	}
}

func TestAgeServiceSuite(t *testing.T) {
	suite.Run(t, &AgeServiceSuite{})
}
