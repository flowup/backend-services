package file

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MockDao is Mock of fileDao object
type MockDao struct {
	meta      []Meta
	idCounter uint
}

// Create will append metadata to server
func (d *MockDao) Create(m *Meta) {
	m.ID = d.idCounter
	d.idCounter++
	m.CreatedAt = time.Now()

	d.meta = append(d.meta, *m)
}

// ReadByID will return metadata of file by given ID
func (d *MockDao) ReadByID(id uint) *Meta {
	for _, v := range d.meta {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

// DeleteByID will remove metadata of file by given ID
func (d *MockDao) DeleteByID(id uint) {
	for i, v := range d.meta {
		if v.ID == id {
			d.meta = append(d.meta[:i], d.meta[i+1:]...)
		}
	}
}

// ReadByHash will return metadata of file by given ID
func (d *MockDao) ReadByHash(hash string) []Meta {
	meta := []Meta{}
	for _, v := range d.meta {
		if v.Hash == hash {
			meta = append(meta, v)
		}
	}
	return meta
}

// Read will return all metadata
func (d *MockDao) Read(m *Meta) []Meta {
	return d.meta
}

// GridSuite defines suite
type GridSuite struct {
	suite.Suite
}

// TestNewGrid tests if Grid service is created correctly
func (s *GridSuite) TestNewGrid() {
	service := NewGrid(&MockDao{})
	assert.NotEqual(s.T(), nil, service)
}

// TestUpload tests if two files are uploaded correctly
func (s *GridSuite) TestUpload() {
	service := NewGrid(&MockDao{})

	file, _ := os.Open("./test_fixtures/test_file_1.txt")
	m := service.Upload(file, file.Name())

	assert.Equal(s.T(), m.Name, file.Name())

	file, _ = os.Open("./test_fixtures/test_file_2.txt")
	m = service.Upload(file, file.Name())

	assert.Equal(s.T(), m.Name, file.Name())

	assert.Equal(s.T(), 2, len(service.dao.Read(&Meta{})))
}

// TestDownload tests if file is correctly downloaded
func (s *GridSuite) TestDownload() {
	service := NewGrid(&MockDao{})

	f, _ := os.Open("./test_fixtures/test_file_1.txt")
	defer f.Close()
	service.Upload(f, f.Name())

	f.Seek(0, 0)
	cmp2, _ := ioutil.ReadAll(f)

	file, _ := os.Open("./test_fixtures/test_file_2.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	file = service.Download(0)

	cmp1, _ := ioutil.ReadAll(file)

	assert.Equal(s.T(), string(cmp2), string(cmp1))
}

// TestGetMeta tests if returned metadata are correct
func (s *GridSuite) TestGetMeta() {
	service := NewGrid(&MockDao{})

	file, _ := os.Open("./test_fixtures/test_file_1.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	file, _ = os.Open("./test_fixtures/test_file_2.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	meta := service.GetMeta(1)

	assert.Equal(s.T(), uint(1), meta.Meta.ID)
	assert.Equal(s.T(), file.Name(), meta.Meta.Name)
}

// TestDelete tests if file is deleted
func (s *GridSuite) TestDelete() {
	service := NewGrid(&MockDao{})

	file, _ := os.Open("./test_fixtures/test_file_1.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	file, _ = os.Open("./test_fixtures/test_file_2.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	service.Delete(0)

	assert.Equal(s.T(), 1, len(service.dao.Read(&Meta{})))
}

// TestGetMetaByHash tests if returned metadata are correct
func (s *GridSuite) TestGetMetaByHash() {
	service := NewGrid(&MockDao{})

	file, _ := os.Open("./test_fixtures/test_file_1.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	file, _ = os.Open("./test_fixtures/test_file_2.txt")
	defer file.Close()
	hash := service.Upload(file, file.Name())

	meta := service.GetMetaByHash(hash.Hash)

	assert.Equal(s.T(), uint(1), meta.Meta.ID)
	assert.Equal(s.T(), file.Name(), meta.Meta.Name)
}

// TestDownloadByHash tests if file is correctly downloaded
func (s *GridSuite) TestDownloadByHash() {
	service := NewGrid(&MockDao{})

	f, _ := os.Open("./test_fixtures/test_file_1.txt")
	defer f.Close()
	hash := service.Upload(f, f.Name())

	f.Seek(0, 0)
	cmp2, _ := ioutil.ReadAll(f)

	file, _ := os.Open("./test_fixtures/test_file_2.txt")
	defer file.Close()
	service.Upload(file, file.Name())

	file = service.DownloadByHash(hash.Hash)

	cmp1, _ := ioutil.ReadAll(file)

	assert.Equal(s.T(), string(cmp2), string(cmp1))
}

// TearDownSuite will run cleanup uploads
func (s *GridSuite) TearDownSuite() {
	os.RemoveAll("./uploads/")
	os.MkdirAll("./uploads/", 0777)
}

//	Passing of our suite (GridSuite) to suite.Run to be able to run the tests
func TestGridSuite(t *testing.T) {
	suite.Run(t, &GridSuite{})
}
