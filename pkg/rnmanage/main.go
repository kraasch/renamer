
package rnmanage

import (
  // local packages.
  auto "github.com/kraasch/renamer/pkg/autorn"
  ctor "github.com/kraasch/renamer/pkg/configurator"
  fsmg "github.com/kraasch/renamer/pkg/fsmanage"

  // external packages.
  "github.com/spf13/afero"

  // standard packages.
  "fmt"
  "bytes"
  "strings"
)

var (
  theFs      afero.Fs
  theInput   string
  conversion string
)

func ConvertByPathList(fs afero.Fs, workDir, conversion, input string) string {
  fmt.Println() // TODO: remove this line later.
  return "Have to implement" // TODO: test and implement.
}

func ConvertByRule(fs afero.Fs, workDir, ruleString, input string) string {
  fmt.Println() // TODO: remove this line later.
  return "Have to implement" // TODO: test and implement.
}

func ConvertByProfile(fs afero.Fs, workDir, configPath, profileName, input string) string {
  if configPath == "" {
    configPath = ctor.PathToDefaultConfig()
  }
  theFs    = fs
  theInput = input
  // open raw content.
  rawToml := ctor.ReadConfig(configPath)
  // parse TOML and apply defined profiles.
  var a auto.AutoRenamer
  a.Parse(rawToml)
  conversion = a.ConvertWith(workDir, profileName, input, fs)
  // return result.
  return ""
}

func ExecuteByValidating() bool {
  return false // TODO: test and implement.
}

func ExecuteByFormatting() string {
  var buf bytes.Buffer
  iis := strings.Split(theInput, "\n")
  ccs := strings.Split(conversion, "\n")
  maxLen := 0
  for _, input := range iis {
    l := len(input)
    if l > maxLen {
      maxLen = l
    }
  }
  for i, input := range iis {
    conv := ccs[i]
    str := fmt.Sprintf("%-*s => %s\n", maxLen, input, conv)
    buf.WriteString(str)
  }
  return buf.String()
}

func ExecuteByApplying() string {
  // apply renames to file system.
  fsmg.DirRename(theFs, theInput, conversion)
  // return result.
  return ""
}

