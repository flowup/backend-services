package filecache

import (
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
  "testing"
)

type ManagerSuite struct {
  suite.Suite

  testFile string
}

func (s *ManagerSuite)SetupSuite() {
  s.testFile = "/run/media/walter/Data/Work/Go/src/flowdock.eu/services/filecache/test_fixtures/template_test"
}

func (s *ManagerSuite) TestCachingTemplates(){
  readData := Cache.LoadFile(s.testFile)
  assert.Equal(s.T(), "THIS IS A TEST TEMPLATE\n", string(readData))
}

func TestManagerSuite(t *testing.T) {
  suite.Run(t, &ManagerSuite{})
}

