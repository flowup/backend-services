package filecache

type FileCacheService interface {
  LoadFile(path string) []byte
  LoadFileNoCache(path string) []byte
  UpdateFile(path string) []byte
  IsFileInCache(path string) bool
}
