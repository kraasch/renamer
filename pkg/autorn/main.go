
package autorn

import (
  "os"
  "bytes"
  "strings"
  // local packages.
  pro "github.com/kraasch/renamer/pkg/profiler"
)

var (
  // NL = fmt.Sprintln() // TODO: use this if it works for all OS'es.
  PS = string(os.PathSeparator)
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
  // apply profile line by line.
  var buf bytes.Buffer
  lines := strings.Split(targetString, "\n")
  for i, line := range lines {
    // split line into path part and file name part.
    lastIndex := strings.LastIndex(line, PS)
    if lastIndex != -1 { // has path separator.
      // apply profile to file name only.
      path     := line[:lastIndex]
      fileName := line[lastIndex+1:]
      buf.WriteString(path + PS + profile.Apply(fileName))
    } else { // has no path separator.
      // apply profile to full line.
      buf.WriteString(profile.Apply(line))
    }
    if i < len(lines) - 1 { // no line break for the last line.
      buf.WriteString("\n")
    }
  }
  // return result.
  return buf.String()
}


