
package dir

import (
  "io/fs"
  "strings"
  "sort"
)

func DirList(fileSystem fs.FS) (out string) {
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

func DirListTree(fileSystem fs.FS) (out string) {
  var fileList []string
  err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
    if err != nil {
      return err // TODO: handle errors appropriately.
    }
    if path == "." {
      return nil // TODO: skip the root directory itself.
    }
    if d.IsDir() {
      fileList = append(fileList, path + "/") // add trailing slash for directories.
    } else {
      fileList = append(fileList, path)
    }
    return nil
  })
  if err != nil {
      return "" // TODO: handle errors appropriately.
  }
  sort.Strings(fileList) // sort alphabetically.
  out = strings.Join(fileList, "\n")
  return out
}

