
package configurator

import (
  "os"
  "path/filepath"
  "strings"
  "errors"
)

const (
  DIRSPERM = 0755
  FILEPERM = 0644
)

var (
  root = "~" // TODO: read the home directory from operating system.
)

type Configurator struct {
  ConfigFileName string
  PathToConfig   string
  DefaultConfig  string
}

func (c *Configurator) AutoReadConfig() string {
  if !c.ExistsDefaultConfig() {
    c.CreateDefaultConfig()
  }
  return c.ReadConfig()
}

func (c *Configurator) SetRoot(path string) {
  root = path
}

func (c *Configurator) ExistsDefaultConfig() bool {
  return FileExists(root + "/" + c.PathToConfig)
}

func (c *Configurator) CreateDefaultConfig() { // TODO: return error.
  CreateFile(root + "/" + c.PathToConfig, c.ConfigFileName, c.DefaultConfig)
}

func (c *Configurator) ReadConfig() string {
  return ReadConfig(root + "/" + c.PathToConfig)
}

func PathToDefaultConfig() string {
  return "Have to implement" // TODO: test and implement.
}

func FileExists(path string) bool {
  if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
    return false
  } else if err == nil {
    return true
  } else {
    // TODO: return error.
    return false
  }
}

func ReadConfig(configPath string) string {
  dat, err := os.ReadFile(configPath)
  if err != nil {
    {} // TODO: report failure.
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

