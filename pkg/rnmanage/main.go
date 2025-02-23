
package rnmanage

import (
  // local packages.
  auto "github.com/kraasch/renamer/pkg/autorn"
  ctor "github.com/kraasch/renamer/pkg/configurator"

  // other packages.
  // "fmt"
)

func Command(configPath, commandType, profileName, input string) string {
  // open raw content.
  rawToml := ctor.ReadConfig(configPath)
  // parse TOML and apply defined profiles.
  var a auto.AutoRenamer
  a.Parse(rawToml)
  output := a.ConvertWith(profileName, input)
  // return result.
  return output
}

