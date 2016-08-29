package timeTracker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// BcryptSuite defines suite
type TrackerSuite struct {
	suite.Suite
	service *Tracker
}

// TestNewBcrypt tests if Bcrypt service is created correctly
func (s *TrackerSuite) SetupTest() {
	s.service = NewTracker()
	assert.NotEqual(s.T(), nil, s.service)
}

// Testing StartingNewEvent if Event is created correctly
func (s *TrackerSuite) TestStartNewEvent() {

	event := s.service.StartNew()
	assert.NotEqual(s.T(), nil, event)
}

// TestStopNil tests if user is not able to put nil event to the Stop() function.
func (s *TrackerSuite) TestStopNil() {

	assert.Equal(s.T(), ErrorNil, s.service.Stop(nil))
}

// TestStartNil tests if user is not able to put nil event to the Start() function.
func (s *TrackerSuite) TestStartNil() {

	assert.Equal(s.T(), ErrorNil, s.service.Start(nil))
}

// TestStartNotStopped tests if user is not able to start existing event if this event was not stopped before
func (s *TrackerSuite) TestStartNotStopped() {

	event := s.service.StartNew()
	assert.Equal(s.T(), ErrorStart, s.service.Start(event))
}

func (s *TrackerSuite) TestTotalCalc() {

	event := s.service.StartNew()
	time.Sleep(3 * time.Second)
	s.service.Stop(event)

	time.Sleep(1 * time.Second)

	s.service.Start(event)
	time.Sleep(4 * time.Second)
	s.service.Stop(event)

	time.Sleep(2 * time.Second)

	s.service.Start(event)
	time.Sleep(1 * time.Second)
	s.service.Stop(event)

	period := (time.Duration(s.service.Total(event)))

	assert.Equal(s.T(), 8, int(period.Seconds()))
}

//	Passing of our suite (TrackerSuite) to suite.Run to be able to run the tests
func TestTrackerSuite(t *testing.T) {
	suite.Run(t, &TrackerSuite{})
}
