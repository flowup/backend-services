package filecache

import (
  "os"
  "runtime"
  "strings"
  "io/ioutil"
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

// LoadFileNoCache will load a template and NOT save it in
// cache.
func (m *NativeCache) LoadFileNoCache(path string) []byte {
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
func (m *NativeCache) LoadFile(path string) []byte{
  if bytes, ok := m.loadedTemplates[path]; ok == true {
    return bytes
  }
  m.loadedTemplates[path] = m.LoadFileNoCache(path)
  return m.loadedTemplates[path]
}