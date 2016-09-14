package filecache

import (
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
		"/src/flowdock.eu/flowup/services/filecache/test_fixtures/template_test"
	s.testFileUpdated = os.Getenv("GOPATH") +
		"/src/flowdock.eu/flowup/services/filecache/test_fixtures/template_test_updated"
}

func (s *ManagerSuite) TestCachingTemplates() {
	readData := Cache.LoadFile(s.testFile)
	assert.Equal(s.T(), "THIS IS A TEST TEMPLATE NONE UPDATED\n", string(readData))
}

func (s *ManagerSuite) TestUpdatingTemplates() {
	readData := Cache.LoadFile(s.testFile)
	assert.Equal(s.T(), "THIS IS A TEST TEMPLATE NONE UPDATED\n", string(readData))

	// Opening file and changing it
	var file, err = os.OpenFile(s.testFile, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("THIS IS A TEST TEMPLATE YES UPDATED!\n")
	if err != nil {
		panic(err)
	}

	// Testing to LoadFile
	readData = Cache.LoadFile(s.testFile)
	assert.NotEqual(s.T(), "THIS IS A TEST TEMPLATE YES UPDATED!\n", string(readData))

	// Testing of update function
	readData = Cache.UpdateTemplate(s.testFile)
	assert.Equal(s.T(), "THIS IS A TEST TEMPLATE YES UPDATED!\n", string(readData))

	// Changing the Test Template to NONE UPDATED
	_, err = file.WriteAt([]byte("THIS IS A TEST TEMPLATE NONE UPDATED"), 0)
	if err != nil {
		panic(err)
	}
	// save changes
	err = file.Sync()
}

func TestManagerSuite(t *testing.T) {
	suite.Run(t, &ManagerSuite{})
}
