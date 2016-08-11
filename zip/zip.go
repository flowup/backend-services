package zip

import "io"

// Service is the interface of zippingService implementation
type Service interface {
	Compress(reader io.Reader, writer io.Writer) error
	Decompress(reader io.Reader, writer io.Writer) error
}

// FileService is the interface of zippingService implementation
type FileService interface {
	CompressFile(source, target string) ([]byte, error)
	DecompressFile(source, target string) ([]byte, error)
}
