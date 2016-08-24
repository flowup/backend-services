package file

import (
	"io"
)

// Service is the interface of fileService implementation
type Service interface {
	Upload(io.ReadWriter)
	Download()
	GetMeta()
}
