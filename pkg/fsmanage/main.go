
package fsmanage

import (
  "fmt"
  "strings"
  "github.com/spf13/afero"
)

func DirRename(fileSystem afero.Fs, originalNames string, targetNames string) {
  fmt.Println("TOAST ########################## ")
  fmt.Println(originalNames)
  fmt.Println("TOAST ########################## ")
  fmt.Println(targetNames)
  fmt.Println("TOAST ########################## ")
  origs  := strings.Split(originalNames, "\n")
  tagets := strings.Split(targetNames, "\n")
  for i := range origs {
    o := origs[i]
    t := tagets[i]
    err := fileSystem.Rename(o, t)
    fmt.Println(err)
  }
}

