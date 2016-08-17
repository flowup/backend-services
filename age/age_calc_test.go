package age

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//	Array of testing dates that are used in TestCalculateAge
var testingDates = [3]time.Time{
	time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2013, 8, 2, 0, 0, 0, 0, time.UTC),
}

//	Defining the suite
type AgeServiceSuite struct {
	suite.Suite
}

//	TestNewAgeService is testing if AgeService is created correctly
func (s *AgeServiceSuite) TestNewAgeService() {
	service := NewAgeService()
	assert.NotEqual(s.T(), nil, service)
}

//	TestCalculateAge is testing calculating algorithm.
func (s *AgeServiceSuite) TestCalculateAge() {
	service := NewAgeService()
	timeNow := time.Now()

	for _, testDate := range testingDates {
		assert.Equal(s.T(), timeNow.Sub(testDate), service.CalculateAge(testDate, timeNow))
	}
}

//	Passing of our suite (AgeServiceSuite) to suite.Run to be able to run the tests
func TestAgeServiceSuite(t *testing.T) {
	suite.Run(t, &AgeServiceSuite{})
}
