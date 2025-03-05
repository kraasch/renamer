
package autorn

import (
  "os"
  "bytes"
  "strings"
  "time"
  "path/filepath"
  // local packages.
  pro "github.com/kraasch/renamer/pkg/profiler"
  // external packages.
  "github.com/spf13/afero"
  "fmt"
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

func normalizePath(path string) string {

  // TODO:work with the absoluate path,
  // if wrapping the filesystem in afero.Fs permits it.
  // // get the absolute path.
  // absPath, err := filepath.Abs(p)
  // if err != nil {
  //   fmt.Printf("Error getting absolute path for %q: %v\n", path, err)
  //   continue
  // }
  // normalizedPath := filepath.Clean(absPath)
  // fmt.Printf("Original: %s, Absolute: %s, Normalized: %s\n", path, absPath, normalizedPath)

  // normalize the path.
  normalizedPath := filepath.Clean(path)
  return normalizedPath
}

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
    // split line into path part and file name part.
    lastIndex := strings.LastIndex(line, PS)
    normalizedPathWd   := normalizePath(workDir)
    normalizedPathLine := normalizePath(line)
    isPrefixed := strings.HasPrefix(normalizedPathLine, normalizedPathWd)
    isRoot     := normalizedPathWd == "."
    if lastIndex != -1 { // has path separator.
      // TODO: evaluate workDir path, in order to also allow paths like:
      // - [ ] "./abc/.."
      // - [ ] "./"
      // - [ ] "."
      if isRoot || isPrefixed {
        // apply profile to file name only.
        path     := line[:lastIndex]
        fileName := line[lastIndex+1:]
        buf.WriteString(path + PS + profile.Apply(fileName, metaInfo))
      } else {
        fmt.Printf("workdir: '%s', line: '%s'\n", normalizedPathWd, normalizedPathLine)
        buf.WriteString(line)
      }
    } else { // has no path separator.
      // do only apply rule to root files if the workdir path is the root.
      if isRoot {
        // apply profile to full line.
        buf.WriteString(profile.Apply(line, metaInfo))
      } else {
        // copy full line.
        buf.WriteString(line)
      }
    }
    if i < len(lines) - 1 { // no line break for the last line.
      buf.WriteString("\n")
    }
  }
  // return result.
  return buf.String()
}


