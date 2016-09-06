package file

import (
	"github.com/jinzhu/gorm"
)

// Meta is model of files uploaded to server
type Meta struct {
	gorm.Model
	Name string
	Hash string
	URL  string
	Type string
}
