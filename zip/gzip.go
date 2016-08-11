package zip

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/gzip"
)

//Gzip is a struct
type Gzip struct {
}

// NewGzip is a constructor creating new Gzip
func NewGzip() *Gzip {
	return &Gzip{}
}

/*CompressFile is compresing func that compress file. It will cut off extension
and zips the file using "github.com/klauspost/compress/gzip" zipping functions */
func (s *Gzip) CompressFile(source, target string) ([]byte, error) {

	//	Reading file and error handle
	reader, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//	Creating target file in target destination
	filename := filepath.Base(source)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	target = filepath.Join(target, fmt.Sprintf("%s.gz", name))
	writer, err := os.Create(target)
	if err != nil {
		return nil, err
	}
	defer writer.Close()

	//	Gzip of the given file
	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	return nil, err
}

//DecompressFile is func decompressing zipped files.
func (s *Gzip) DecompressFile(source, target string) error {

	//	Reading file and error handle
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	//	Reading zipped files
	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	//	Creating unzipped file
	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}

//Compress is compresing data from io.Reader and sending them to io.Writer
func (s *Gzip) Compress(reader io.Reader, writer io.Writer) error {

	archiver := gzip.NewWriter(writer)
	defer archiver.Close()

	_, err := io.Copy(archiver, reader)
	return err
}

//Decompress is compresing data from  and sending them to io.Writer
func (s *Gzip) Decompress(reader io.Reader, writer io.Writer) error {

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	_, err = io.Copy(writer, archive)
	return err

}
