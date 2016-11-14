package filecache

import (
  "fmt"
  "os"
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type ManagerSuite struct {
  suite.Suite

  testFile        string
  testFileUpdated string
}

func (s *ManagerSuite) SetupSuite() {
  s.testFile = os.Getenv("GOPATH") +
    "/src/github.com/flowup/backend-services/filecache/test_fixtures/template_test"
  s.testFileUpdated = os.Getenv("GOPATH") +
    "/src/github.com/flowup/backend-services/filecache/test_fixtures/template_test_updated"
}

func (s *ManagerSuite) TestCachingTemplates() {
  readData := Cache.LoadFile(s.testFile)
  assert.Equal(s.T(), "THIS IS A TEST TEMPLATE\n", string(readData))
}

func (s *ManagerSuite) TestUpdatingTemplates() {
  readData := Cache.LoadFile(s.testFile)
  assert.Equal(s.T(), "THIS IS A TEST TEMPLATE\n", string(readData))

  os.Truncate(s.testFile, 0)

  // Opening file and changing it
  var file, err = os.OpenFile(s.testFile, os.O_RDWR, 0644)
  assert.Nil(s.T(), err)
  defer file.Close()

  _, err = file.WriteString("UPDATED TEST TEMPLATE!\n")
  assert.Nil(s.T(), err)

  // Testing to LoadFile
  readData = Cache.LoadFile(s.testFile)
  fmt.Println(string(readData))
  assert.NotEqual(s.T(), "UPDATED TEST TEMPLATE!\n", string(readData))

  // Testing of update function
  readData = Cache.UpdateFile(s.testFile)
  fmt.Println(string(readData))
  assert.Equal(s.T(), "UPDATED TEST TEMPLATE!\n", string(readData))

  // Changing the Test Template to NONE UPDATED

  _, err = file.WriteAt([]byte("THIS IS A TEST TEMPLATE\n"), 0)

  assert.Nil(s.T(), err)
  // save changes
  err = file.Sync()

}

func TestManagerSuite(t *testing.T) {
  suite.Run(t, &ManagerSuite{})
}
