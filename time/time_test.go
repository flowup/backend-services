package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TimeSuite defines suite
type TimeSuite struct {
	suite.Suite
}

// TestNewMock tests if mock time is created correctly
func (s *TimeSuite) TestNewTime() {
	mock := NewMockTime(time.Now())

	assert.NotEqual(s.T(), nil, mock)
}

// TestNow tests if mock time return correct time
func (s *TimeSuite) TestNow() {
	now := time.Now()
	mock := NewMockTime(now)

	assert.Equal(s.T(), now, mock.Now())
}

// TestSetTime tests if mock time sets time correctly
func (s *TimeSuite) TestSetTime() {
	now := time.Now()
	mock := NewMockTime(now)

	assert.Equal(s.T(), now, mock.Now())

	zero := time.Time{}
	mock.SetTime(zero)

	assert.Equal(s.T(), zero, mock.Now())
}

// TestSleep tests if mock time will increment time correctly
func (s *TimeSuite) TestSleep() {
	now := time.Now()
	mock := NewMockTime(now)

	mock.Sleep(10 * time.Hour)

	assert.Equal(s.T(), now.Add(10*time.Hour), mock.Now())
}

// TestSleep tests if mock time will rollback time correctly
func (s *TimeSuite) TestRollback() {
	now := time.Now()
	mock := NewMockTime(now)

	mock.Rollback(10 * time.Hour)

	assert.Equal(s.T(), now.Add(-10*time.Hour), mock.Now())
}

//	Passing of our suite (TimeSuite) to suite.Run to be able to run the tests
func TestGridSuite(t *testing.T) {
	suite.Run(t, &TimeSuite{})
}
