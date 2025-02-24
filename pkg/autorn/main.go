
package autorn

import (
  "os"
  "bytes"
  "strings"
  "time"
  // local packages.
  pro "github.com/kraasch/renamer/pkg/profiler"
  // external packages.
  "github.com/spf13/afero"
)

type FileInfo struct {
  // empty.
}

func (ti FileInfo) CurrentDate() string {
  current_time := time.Now().Local()
  return current_time.Format("2006-01-02")
}

func (ti FileInfo) CreationDate() string {
  return "2020-12-20" // TODO: implement.
}

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

func (a *AutoRenamer) ConvertWith(profileName, targetString string, fs afero.Fs) string {
  // TODO: implement FileInfo.
  // - for [id^], [id.] and [id$] add the CurrentDate().
  // - for [ic^], [ic.] and [ic$] add the CreationDate() from the file system.
  metaInfo := FileInfo{}
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
      buf.WriteString(path + PS + profile.Apply(fileName, metaInfo))
    } else { // has no path separator.
      // apply profile to full line.
      buf.WriteString(profile.Apply(line, metaInfo))
    }
    if i < len(lines) - 1 { // no line break for the last line.
      buf.WriteString("\n")
    }
  }
  // return result.
  return buf.String()
}


