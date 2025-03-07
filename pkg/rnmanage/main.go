
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
  theFs         afero.Fs
  theInput      string
  theConversion string
)

func ConvertByPathList(fs afero.Fs, workDir, conv, input string) string {
  theFs         = fs
  theInput      = input
  theConversion = conv
  return "" // TODO: remove return.
}

func ConvertByRule(fs afero.Fs, workDir, ruleString, input string) string {
  theFs         = fs
  theInput      = input
  theConversion = auto.ConvertWithRule(workDir, ruleString, input, fs)
  return "" // TODO: remove return.
}

func ConvertByProfile(fs afero.Fs, workDir, configPath, profileName, input string) string {
  theFs    = fs
  theInput = input
  // open raw content.
  // TODO: use this here:
  // c := Configurator{ ... }
  // c.AutoReadConfig()
  rawToml := ctor.ReadConfig(configPath)
  // parse TOML and apply defined profiles.
  var a auto.AutoRenamer
  a.Parse(rawToml)
  theConversion = a.ConvertWith(workDir, profileName, input, fs)
  // return result.
  return "" // TODO: remove return.
}

func ExecuteByValidating() bool {
  allEqual := true
  iis := strings.Split(theInput, "\n")
  ccs := strings.Split(theConversion, "\n")
  for i, input := range iis {
    conv := ccs[i]
    if conv != input {
      allEqual = false
      break
    }
  }
  return allEqual
}

func ExecuteByFormatting() string {
  var buf bytes.Buffer
  iis := strings.Split(theInput, "\n")
  ccs := strings.Split(theConversion, "\n")
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
  fsmg.DirRename(theFs, theInput, theConversion)
  // return result.
  return ""
}

