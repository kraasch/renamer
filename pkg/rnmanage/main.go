
package rnmanage

import (
  // local packages.
  auto "github.com/kraasch/renamer/pkg/autorn"
  ctor "github.com/kraasch/renamer/pkg/configurator"
  fsmg "github.com/kraasch/renamer/pkg/fsmanage"

  // external packages.
  "github.com/spf13/afero"

  // standard packages.
  "bytes"
  "strings"
  "fmt"
)

var (
  theFs         afero.Fs
  theInput      string
  theConversion string
  NL = fmt.Sprintln()
  DEFAULT_CONFIG_NAME = "renamer.toml"
  DEFAULT_CONFIG_PATH = ".config/renamer/"
  DEFAULT_CONFIG_TEXT =
    `title = "Basic Conf"                     ` + NL +
    `[profiles]                               ` + NL +
    `  [profiles.lcase-txt]                   ` + NL +
    `    name = "LowerCase"                   ` + NL +
    `    [profiles.lcase-txt.profile_rule]    ` + NL +
    `      word_separators = " ()"            ` + NL +
    `      delete_chars    = ""               ` + NL +
    `      small_gap_mark  = "-"              ` + NL +
    `      big_gap_mark    = "_"              ` + NL +
    `      conversions     = "cAa"            ` + NL +
    `      modes_string    = ""               ` + NL +
    `  [profiles.prettify-txt]                ` + NL +
    `    name = "SomeProfile"                 ` + NL +
    `    [profiles.prettify-txt.profile_rule] ` + NL +
    `      word_separators = " ()"            ` + NL +
    `      delete_chars    = ""               ` + NL +
    `      small_gap_mark  = "-"              ` + NL +
    `      big_gap_mark    = "_"              ` + NL +
    `      conversions     = "caA"            ` + NL +
    `      modes_string    = ""               ` + NL +
    `                                         `
  CONFIG = ctor.Configurator{
    ConfigFileName: DEFAULT_CONFIG_NAME,
    PathToConfig:   DEFAULT_CONFIG_PATH,
    DefaultConfig:  DEFAULT_CONFIG_TEXT,
  }
)

func SetTestRoot(newRoot string) string  {
  root := CONFIG.Root()
  CONFIG.SetRoot(newRoot)
  return root
}

func SetTestConfig(newRoot, configPath, configName string) string  {
  CONFIG.PathToConfig = configPath
  CONFIG.ConfigFileName = configName
  return SetTestRoot(newRoot)
}

func ListProfiles() string {
  rawToml := CONFIG.ReadConfig()
  var a auto.AutoRenamer
  a.Parse(rawToml)
  return a.ListProfiles()
}

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
  rawToml := CONFIG.AutoReadConfig()
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

