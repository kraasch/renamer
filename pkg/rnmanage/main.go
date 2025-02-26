
package rnmanage

import (
  // local packages.
  auto "github.com/kraasch/renamer/pkg/autorn"
  ctor "github.com/kraasch/renamer/pkg/configurator"
  fsmg "github.com/kraasch/renamer/pkg/fsmanage"

  // external packages.
  "github.com/spf13/afero"
  "fmt"
)

var (
  theFs      afero.Fs
  theInput   string
  conversion string
)

func ConvertByPathList(fs afero.Fs, conversion, input string) string {
  fmt.Println() // TODO: remove this line later.
  return "Have to implement" // TODO: test and implement.
}

func ConvertByRule(fs afero.Fs, ruleString, input string) string {
  fmt.Println() // TODO: remove this line later.
  return "Have to implement" // TODO: test and implement.
}

func AutoConvertByProfile(fs afero.Fs, configPath, profileName, input string) string {
  path := configPath
  if path == "" {
    path = ctor.PathToDefaultConfig()
  }
  output := ConvertByProfile(fs, path, profileName, input)
  return output
}

func ConvertByProfile(fs afero.Fs, configPath, profileName, input string) string {
  theFs    = fs
  theInput = input
  // open raw content.
  rawToml := ctor.ReadConfig(configPath)
  // parse TOML and apply defined profiles.
  var a auto.AutoRenamer
  a.Parse(rawToml)
  conversion = a.ConvertWith(profileName, input, fs)
  // return result.
  return ""
}

func ExecuteByValidating() string {
  return "Have to implement" // TODO: test and implement.
}

func ExecuteByPrinting() string {
  return "Have to implement" // TODO: test and implement.
}

func ExecuteByApplying() string {
  // apply renames to file system.
  fsmg.DirRename(theFs, theInput, conversion)
  // return result.
  return ""
}

