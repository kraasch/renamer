
package fsmanage

import (
  "strings"
  "github.com/spf13/afero"
)

func DirRename(fileSystem afero.Fs, originalNames string, targetNames string) {
  origs  := strings.Split(originalNames, "\n")
  tagets := strings.Split(targetNames, "\n")
  for i := range origs {
    o := origs[i]
    t := tagets[i]
    _ = fileSystem.Rename(o, t)
  }
}

