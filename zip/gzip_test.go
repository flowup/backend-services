package zip

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ZipServiceSuite struct {
	suite.Suite
	service *ZipService
}

//SetupTest
func (s *ZipServiceSuite) SetupTest() {
	s.service = NewZipService()
	assert.NotEqual(s.T(), nil, s.service)
}

/*TestCompressDecompressFile is testing funcs CompressFile and DecompressFile
it takes file test.txt from test_fixtures and zips and after unzips it. */
func (s *ZipServiceSuite) TestCompressDecompressFile() {

	s.service.CompressFile("test_fixtures/test.txt", "test_fixtures")
	s.service.DecompressFile("test_fixtures/test.gz", "/tmp")

	dat, err := ioutil.ReadFile("/tmp/test.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove("test_fixtures/test.gz")

	assert.Equal(s.T(), "Test string! \n", string(dat))
}

//TestCompressDecompress testing basic compressing of bytes coming from io.Reader
func (s *ZipServiceSuite) TestCompressDecompress() {
	//	Source buffer simulating data in io.Reader
	sender := bytes.NewBuffer([]byte("Hello World"))

	//	Two buffers. One for output of zipping func and second for output of unzipping
	receiver := bytes.NewBuffer(nil)
	receiver2 := bytes.NewBuffer(nil)

	//	Compressing and decompressing
	s.service.Compress(sender, receiver)
	s.service.Decompress(receiver, receiver2)

	//	Testing if testing string is the same as before zipping
	assert.Equal(s.T(), "Hello World", string(receiver2.Bytes()))
}

//TestZipServiceSuite is func that starts testing
func TestZipServiceSuite(t *testing.T) {
	suite.Run(t, &ZipServiceSuite{})
}
