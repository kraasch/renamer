
package autorn

import (
  // standard packages.
  "os"
  "bytes"
  "strings"
  "time"
  "path/filepath"
  "fmt"

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
  NLine = fmt.Sprintln() // TODO: use this if it works for all OS'es.
  PS    = string(os.PathSeparator)
)

type AutoRenamer struct {
  config pro.Config
}

func (a *AutoRenamer) Config() pro.Config {
  return a.config
}

func (a *AutoRenamer) Parse(toml string) {
  cfg := pro.ReadRawProfileConfig(toml)
  a.config = cfg
}

func normalizePath(path string) string {
  normalizedPath := filepath.Clean(path) // normalize the path.
  return normalizedPath
}

func convertLine(line, workDir string, profile *pro.Profile, metaInfo FileInfo) string {
  // split line into path part and file name part.
  lastIndex  := strings.LastIndex(line, PS)
  normWdir   := normalizePath(workDir)
  normLine   := normalizePath(line)
  isPrefixed := strings.HasPrefix(normLine, normWdir)
  isRoot     := normWdir == "."
  hasSep     := lastIndex != -1
  if hasSep && (isRoot || isPrefixed) {
    // apply profile to file name only.
    path     := line[:lastIndex]
    fileName := line[lastIndex+1:]
    return path + PS + profile.Apply(fileName, metaInfo)
  }
  if !hasSep && isRoot {
    // do only apply rule to root files if the workdir path is the root.
    // apply profile to full line.
    return profile.Apply(line, metaInfo)
  }
  return line
}

// TODO: extract function which operates on profile. (1/2)
func (a *AutoRenamer) ConvertWith(workDir, profileName, targetString string, fs afero.Fs) string {
  // TODO: implement FileInfo.
  // - for [id^], [id.] and [id$] add the CurrentDate().
  // - for [ic^], [ic.] and [ic$] add the CreationDate() from the file system.
  metaInfo := FileInfo{}
  profile := a.config.Profiles[profileName]
  // apply profile line by line.
  var buf bytes.Buffer
  lines := strings.Split(targetString, "\n")
  for i, line := range lines {
    buf.WriteString(convertLine(line, workDir, profile, metaInfo))
    if i < len(lines) - 1 {  // no line break for the last line.
      buf.WriteString(NLine) // break line.
    }
  }
  // return result.
  return buf.String()
}

// TODO: extract function which operates on profile. (2/2)
func ConvertWithRule(workDir, ruleString, targetString string, fs afero.Fs) string {
  profile := pro.ProfileFromRuleString(ruleString)
  var buf bytes.Buffer
  lines := strings.Split(targetString, "\n")
  for i, line := range lines {
    buff := convertLine(line, workDir, &profile, FileInfo{})
    buf.WriteString(buff)
    if i < len(lines) - 1 {  // no line break for the last line.
      buf.WriteString(NLine) // break line.
    }
  }
  // return result.
  return buf.String()
}

