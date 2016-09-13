package filecache

import (
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
  "testing"
  "os"
)

type ManagerSuite struct {
  suite.Suite

  testFile string
}

func (s *ManagerSuite)SetupSuite() {
  s.testFile = os.Getenv("GOPATH") +
    "/src/flowdock.eu/flowup/services/filecache/test_fixtures/template_test"
}

func (s *ManagerSuite) TestCachingTemplates(){
  readData := Cache.LoadFile(s.testFile)
  assert.Equal(s.T(), "THIS IS A TEST TEMPLATE\n", string(readData))
}

func TestManagerSuite(t *testing.T) {
  suite.Run(t, &ManagerSuite{})
}

