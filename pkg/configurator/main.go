
package configurator

import (
  "os"
  "path/filepath"
  "strings"
)

const (
  DIRSPERM = 0755
  FILEPERM = 0644
)

func ReadConfig(configPath string) string {
  dat, err := os.ReadFile(configPath)
  if err != nil {
    {}
    // panic(err)
  }
  return string(dat)
}

func CreateFile(pathToFile, fileName, fileContent string) {
  dirs := strings.Split(pathToFile, "/") // TODO: change for windows.
  dirBuf := ""
  for _, dir := range dirs {
    dirBuf += dir + "/"
    if err := os.MkdirAll(dirBuf, DIRSPERM); err != nil {
      {} // TODO: report failure.
    }
  }
  full := filepath.Join(pathToFile, fileName)
  if err := os.WriteFile(full, []byte(fileContent), FILEPERM); err != nil {
    {} // TODO: report failure.
  }
}

