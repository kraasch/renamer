
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
  root, _ = os.UserHomeDir()
)

type Configurator struct {
  ConfigFileName string
  PathToConfig   string
  DefaultConfig  string
}

func (c *Configurator) getConfigPath() string {
  return root + "/" + c.PathToConfig
}

func (c *Configurator) getFullConfigPath() string {
  return root + "/" + c.PathToConfig + "/" + c.ConfigFileName
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

func (c *Configurator) Root() string {
  return root
}

func (c *Configurator) ExistsDefaultConfig() bool {
  return FileExists(c.getFullConfigPath())
}

func (c *Configurator) CreateDefaultConfig() { // TODO: return error.
  CreateFile(c.getConfigPath(), c.ConfigFileName, c.DefaultConfig)
}

func (c *Configurator) ReadConfig() string {
  return ReadConfig(c.getFullConfigPath())
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

