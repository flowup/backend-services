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
}

// NewGrid creates new Grid
func NewGrid(dao fileDao) *Grid {
	return &Grid{dao}
}

// Upload will save given file to server and return its metadata
func (g *Grid) Upload(file io.ReadWriter, name string) *Meta {
	// Create unique name of file
	now := time.Now().Unix()
	h := sha1.New()
	io.WriteString(h, strconv.FormatInt(now, 10))
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
	meta := &Meta{Name: name, URL: out.Name(), Type: fext}
	g.dao.Create(meta)

	return meta
}

// Download will return file from server by given metadata of this file
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
func (g *Grid) GetMeta(id uint) *Meta {
	return g.dao.ReadByID(id)
}

// Delete will remove data from server
func (g *Grid) Delete(id uint) {
	m := g.dao.ReadByID(id)
	os.Remove(m.URL)
	g.dao.DeleteByID(id)
}
