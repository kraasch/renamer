
package autorn

import (
  // local packages.
  pro "github.com/kraasch/renamer/pkg/profiler"
)

type AutoRenamer struct {
  config pro.Config
}

func (a *AutoRenamer) Parse(toml string) {
  cfg := pro.ReadRawProfileConfig(toml)
  a.config = cfg
}

func (a *AutoRenamer) ConvertWith(profileName, targetString string) string {
  profile := a.config.Profiles[profileName]
  output := profile.Apply(targetString)
  return output
}


