
package rnmanage

import (
  "os"
  pro "github.com/kraasch/renamer/pkg/profiler"
)

func Command(configPath, commandType, profileName, input string) string {
  dat, err := os.ReadFile("." + "/" + configPath)
  if err != nil {
    panic(err)
  }
  cfg := pro.ReadRawProfileConfig(string(dat))
  profile := cfg.Profiles[profileName]
  rule := profile.ProfileRule
  return rule.WordSeparators
}

