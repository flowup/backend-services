package file

import "io"

// Service is the interface of fileService implementation
type Service interface {
	Upload(file io.ReadWriter, name string) *Meta
	Download(id uint) io.ReadWriter
	GetMeta(id uint) *Meta
	Delete(id uint)
}
