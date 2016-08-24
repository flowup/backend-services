package file

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Grid is implementation of fileService
type Grid struct {
}

// MetaData are data of file uploaded to server
type MetaData struct {
	fileURL  string
	fileType string
}

// Upload will save given file to server and return its metadata
func (g *Grid) Upload(f io.ReadWriter, name string) (MetaData, error) {
	meta := MetaData{}

	// Create unique name of file
	now := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(now, 10))
	fname := fmt.Sprintf("%x", h.Sum(nil))

	// Get extension of file
	fext := filepath.Ext(name)
	meta.fileType = fext

	out, err := os.OpenFile("./uploads/"+fname+fext, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return meta, err
	}
	defer out.Close()

	meta.fileURL = out.Name()

	if _, err = io.Copy(out, f); err != nil {
		return meta, err
	}

	return meta, nil
}

// Download will return file from server by given metadata of this file
func (g *Grid) Download(meta MetaData) (*os.File, error) {

	return nil, nil
}

// GetMeta will return metadata of file from server
func (g *Grid) GetMeta() (MetaData, error) {
	meta := MetaData{}

	return meta, nil
}
