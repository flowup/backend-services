package file

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Grid is implementation of fileService
type Grid struct {
	dao fileDao
}

// fileDao is Dao interface of model
type fileDao interface {
	ReadByID(id uint) *Meta
	DeleteByID(id uint)
	Create(m *Meta)
	Read(m *Meta) []Meta
	ReadByHash(hash string) []Meta
}

// NewGrid creates new Grid
func NewGrid(dao fileDao) *Grid {
	return &Grid{dao}
}

// MetaData is type of returned metadata
type MetaData struct {
	Meta *Meta
	Size int64
}

// Upload will save given file to server and return its metadata
func (g *Grid) Upload(file io.Reader, name string) *Meta {
	// Create unique name of file
	now := time.Now()
	h := sha1.New()
	io.WriteString(h, strconv.FormatInt(int64(now.Second()), 10)+strconv.FormatInt(int64(now.Nanosecond()), 10))
	fname := fmt.Sprintf("%x", h.Sum(nil))

	// Get extension of file
	fext := filepath.Ext(name)

	// Create new file with unique name in directory
	out, err := os.OpenFile("./uploads/"+fname+fext, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}
	defer out.Close()

	// Copy data to new file
	if _, err = io.Copy(out, file); err != nil {
		return nil
	}

	// Add metadata to database
	meta := &Meta{Name: name, URL: out.Name(), Type: fext, Hash: fname}
	g.dao.Create(meta)

	return meta
}

// Download will return file from server
func (g *Grid) Download(id uint) *os.File {
	meta := g.dao.ReadByID(id)

	if meta == nil {
		return nil
	}

	file, err := os.Open(meta.URL)
	if err != nil {
		return nil
	}
	return file
}

// GetMeta will return metadata of file from server
func (g *Grid) GetMeta(id uint) MetaData {
	meta := g.dao.ReadByID(id)
	fi, _ := os.Stat(meta.URL)
	return MetaData{Meta: meta, Size: fi.Size()}
}

// Delete will remove data from server
func (g *Grid) Delete(id uint) {
	m := g.dao.ReadByID(id)
	os.Remove(m.URL)
	g.dao.DeleteByID(id)
}

// GetMetaByHash will return metadata of file from server by given hash string
func (g *Grid) GetMetaByHash(hash string) MetaData {
	meta := g.dao.ReadByHash(hash)[0]
	fi, _ := os.Stat(meta.URL)
	return MetaData{Meta: &meta, Size: fi.Size()}
}

// DownloadByHash will return file from server by given hash string
func (g *Grid) DownloadByHash(hash string) *os.File {
	meta := g.dao.ReadByHash(hash)[0]

	file, err := os.Open(meta.URL)
	if err != nil {
		return nil
	}

	return file
}
