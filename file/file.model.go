package file

import (
	"github.com/jinzhu/gorm"
)

// File is model of files uploaded to server
type File struct {
	gorm.Model
	fileID   int64
	fileURL  string
	fileType string
}
