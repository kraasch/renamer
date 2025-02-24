
package rnmanage

import (
  // local packages.
  auto "github.com/kraasch/renamer/pkg/autorn"
  ctor "github.com/kraasch/renamer/pkg/configurator"
  fsmg "github.com/kraasch/renamer/pkg/fsmanage"

  // external packages.
  "github.com/spf13/afero"

  // other packages.
  // "fmt"
)

func Command(fileSystem afero.Fs, configPath, profileName, input string) string {

  // open raw content.
  rawToml := ctor.ReadConfig(configPath)

  // parse TOML and apply defined profiles.
  var a auto.AutoRenamer
  a.Parse(rawToml)
  targetNames := a.ConvertWith(profileName, input, fileSystem)

  // apply renames to file system.
  fsmg.DirRename(fileSystem, input, targetNames)

  // return result.
  return ""
}

