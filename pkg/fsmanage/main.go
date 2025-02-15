
package fsmanage

import (
  "io/fs"
  "strings"
)

func DirRename(fileSystem fs.FS, targetNames string) (out string) {
  dirEntries, err := fs.ReadDir(fileSystem, ".")
  if err != nil {
    return "" // TODO: handle errors appropriately.
  }
  var entries []string
  for _, entry := range dirEntries {
    name := entry.Name()
    if entry.IsDir() {
      name += "/" // add trailing slash for directories.
    }
    entries = append(entries, name)
  }
  out = strings.Join(entries, "\n")
  return out
}

