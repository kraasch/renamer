
package fsmanage

import (
  "io"
  "os"
  "io/fs"
  "strings"
  // "path/filepath"
)

// https://stackoverflow.com/questions/16742331/how-to-mock-abstract-filesystem-in-go
// MOCK FS START.
var fss fileSystem = osFS{}
type fileSystem interface {
    Open(name string) (file, error)
    Stat(name string) (os.FileInfo, error)
}
type file interface {
    io.Closer
    io.Reader
    io.ReaderAt
    io.Seeker
    Stat() (os.FileInfo, error)
}
type osFS struct{} // osFS implements fileSystem using the local disk.
func (osFS) Open(name string) (file, error)        { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
// MOCK FS END.

func DirRename(fileSystem fs.FS, originalNames string, targetNames string) (out string) {
  origs  := strings.Split(originalNames, "\n")
  tagets := strings.Split(targetNames, "\n")
  for i, _ := range origs {
    o := origs[i]
    t := tagets[i]
    _ = fss.Rename(o, t)
  }

  return "TODO: implement." // TODO: implement.
}

