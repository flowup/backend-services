package zip

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GzipSuite struct {
	suite.Suite
	service *Gzip
}

//SetupTest
func (s *GzipSuite) SetupTest() {
	s.service = NewGzip()
	assert.NotEqual(s.T(), nil, s.service)
}

/*TestCompressDecompressFile is testing funcs CompressFile and DecompressFile
it takes file test.txt from test_fixtures and zips and after unzips it. */
func (s *GzipSuite) TestCompressDecompressFile() {

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
func (s *GzipSuite) TestCompressDecompress() {
	const testString = "Hello World"

	//	Source buffer simulating data in io.Reader
	content := bytes.NewBuffer([]byte(testString))

	//	Two buffers. One for output of zipping func and second for output of unzipping
	compressed := bytes.NewBuffer(nil)
	decompressed := bytes.NewBuffer(nil)

	//	Compressing and decompressing
	s.service.Compress(content, compressed)
	s.service.Decompress(compressed, decompressed)

	//	Testing if testing string is the same as before zipping
	assert.Equal(s.T(), testString, string(decompressed.Bytes()))
}

//TestGzipSuite is func that starts testing
func TestGzipSuite(t *testing.T) {
	suite.Run(t, &GzipSuite{})
}
