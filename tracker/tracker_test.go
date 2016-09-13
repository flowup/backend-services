package tracker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// BcryptSuite defines suite
type TimeTrackerSuite struct {
	suite.Suite
	service *TimeTracker
}

// TestNewBcrypt tests if Bcrypt service is created correctly
func (s *TimeTrackerSuite) SetupTest() {
	s.service = NewTimeTracker()
	assert.NotEqual(s.T(), nil, s.service)
}

// Testing StartingNewEvent if Event is created correctly
func (s *TimeTrackerSuite) TestStartNewEvent() {

	event := s.service.StartNew()
	assert.NotEqual(s.T(), nil, event)
}

// TestStopNil tests if user is not able to put nil event to the Stop() function.
func (s *TimeTrackerSuite) TestStopNil() {

	assert.Equal(s.T(), ErrorNil, s.service.Stop(nil))
}

// TestStartNil tests if user is not able to put nil event to the Start() function.
func (s *TimeTrackerSuite) TestStartNil() {

	assert.Equal(s.T(), ErrorNil, s.service.Start(nil))
}

// TestStartNotStopped tests if user is not able to start existing event if this event was not stopped before
func (s *TimeTrackerSuite) TestStartNotStopped() {

	event := s.service.StartNew()
	assert.Equal(s.T(), ErrorStart, s.service.Start(event))
}

func (s *TimeTrackerSuite) TestTotalCalc() {

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

//	Passing of our suite (TimeTrackerSuite) to suite.Run to be able to run the tests
func TestTimeTrackerSuite(t *testing.T) {
	suite.Run(t, &TimeTrackerSuite{})
}
