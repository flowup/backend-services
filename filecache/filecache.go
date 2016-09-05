package filecache

type FileCacheService interface {
  LoadFile(path string) []byte
  LoadFileNoCache(path string) []byte
}
