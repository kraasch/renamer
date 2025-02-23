
package rnmanage

import (
  // local packages.
  pro "github.com/kraasch/renamer/pkg/profiler"
  // autorn "github.com/kraasch/renamer/pkg/autorn"
  // other packages.
  "os"
  // misc packages.
  "fmt"
)

func Command(configPath, commandType, profileName, input string) string {
  dat, err := os.ReadFile("." + "/" + configPath)
  if err != nil {
    panic(err)
  }
  cfg := pro.ReadRawProfileConfig(string(dat))
  profile := cfg.Profiles[profileName]
  fmt.Println("CONFIG:", cfg)

  // parse TOML and apply defined profiles.
  // var auto AutoRenamer
  // auto.Parse(toml)
  // output := auto.ConvertWith(profileName, name)
  // return output
  return fmt.Sprintf("%#v\n", profile)
}

