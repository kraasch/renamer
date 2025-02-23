
package rnmanage

import (
  pro "github.com/kraasch/renamer/pkg/profiler"
  "os"
  "strings"
  "bytes"
)

func Command(configPath, commandType, profileName, input string) string {
  dat, err := os.ReadFile("." + "/" + configPath)
  if err != nil {
    panic(err)
  }
  cfg := pro.ReadRawProfileConfig(string(dat))
  profile := cfg.Profiles[profileName]

  // get results.
  var buf bytes.Buffer
  lines := strings.Split(input, "\n")
  for i, line := range lines {
    if i == len(lines) - 1 && line == "" { // NOTE: remove last line break from pipe.
      break
    }
    buf.WriteString(profile.Apply(line) + "\n")
  }
  return buf.String()
}

