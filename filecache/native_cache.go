package filecache

import (
  "io/ioutil"
  "os"
  "runtime"
  "strings"
  "path/filepath"
)

// instance of a cache
var Cache *NativeCache

// NativeCache is a structure holding cached templates
type NativeCache struct {
  loadedTemplates map[string][]byte
}

func init() {
  Cache = &NativeCache{
    loadedTemplates: make(map[string][]byte),
  }
}

// getAbsolutePath is auxiliary function that will transform
// any relative path to absolute path. If such transformation
// is not possible, panic occurs
func getAbsolutePath(path string) string {
  var err error
  if !filepath.IsAbs(path) {
    path, err = filepath.Abs(path)
    if err != nil {
      panic(err)
    }
  }
  return path
}

// LoadFileNoCache will load a template and NOT save it in
// cache.
func (m *NativeCache) LoadFileNoCache(path string) []byte {
  path = getAbsolutePath(path)

  if runtime.GOOS == "windows" {
    path = strings.Replace(path, "/", "\\", -1)
  }
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }
  return bytes
}

// LoadTemplate will look into cache if the requested
// template is not loaded already, if it is, it will return it
// without any further ado, if it is not, it will load it,
// save it to cache and return it.
func (m *NativeCache) LoadFile(path string) []byte {
  path = getAbsolutePath(path)

  if bytes, ok := m.loadedTemplates[path]; ok == true {
    return bytes
  }
  m.loadedTemplates[path] = m.LoadFileNoCache(path)
  return m.loadedTemplates[path]
}

// UpdateFile is updating existing template in map of templates
func (m *NativeCache) UpdateFile(path string) []byte {
  path = getAbsolutePath(path)

  m.loadedTemplates[path] = m.LoadFileNoCache(path)
  return m.loadedTemplates[path]
}

// IsFileInCache will return true if file is in cache, false otherwise
func (m *NativeCache) IsFileInCache(path string) bool {
  path = getAbsolutePath(path)

  _, ok := m.loadedTemplates[path]
  return ok
}