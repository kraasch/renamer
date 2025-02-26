
package rnmanage

import (
  // unit testing.
  "testing"
  gt "github.com/kraasch/gotest/gotest"

  // cli testing.
  "fmt"
  "os/exec"
  "strings"

  // local packages.
  tu "github.com/kraasch/renamer/pkg/testutil"
)

var (
  NL = fmt.Sprintln()
  BASIC_CONF =
  `title = "Basic Conf"                     ` + NL +
  `[profiles]                               ` + NL +
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
  FIND_TXT =
  "NOTES.txt"             + NL +
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/APPLES.txt"     + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/COCONUTS.txt"   + NL +
  "shapes/"               + NL +
  "shapes/CIRCLE.txt"     + NL +
  "shapes/SQUARE.txt"     + NL +
  "shapes/TRIANGLE.txt"
  LS_TXT =
  "NOTES.txt"             + NL +
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/apples.txt"     + NL +
  "fruits/bananas.txt"    + NL +
  "fruits/coconuts.txt"   + NL +
  "shapes/"               + NL +
  "shapes/circle.txt"     + NL +
  "shapes/square.txt"     + NL +
  "shapes/triangle.txt"
  FIND_ALL =
  "NOTES.txt"             + NL +
  "config/"               + NL +
  "config/GENERAL.config" + NL +
  "fruits/"               + NL +
  "fruits/APPLES.txt"     + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/COCONUTS.txt"   + NL +
  "shapes/"               + NL +
  "shapes/CIRCLE.txt"     + NL +
  "shapes/SQUARE.txt"     + NL +
  "shapes/TRIANGLE.txt"
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

type CommandList []struct {
  Name string
  Args []string
}

func simulatePipe(commands CommandList, path string, t *testing.T) string {
  var output []byte
  for i, c := range commands {
    cmd := exec.Command(c.Name, c.Args...)
    cmd.Dir = path // execute within diretory of test file system.
    cmd.Stdin = strings.NewReader(string(output))
    output, _ = cmd.Output()
    t.Logf("%d > %s %v \t==> %s\n", i, c.Name, c.Args, output)
  }
  return string(output)
}

var suites = []gt.TestSuite{

  /*
  * Test ConvertByProfile()
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      configPath    := in.InputArr[0]
      configName    := in.InputArr[1]
      profileName   := in.InputArr[2]
      configContent := in.InputArr[3]
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      inputListing  := mdir.ListTree()
      conf          := mdir.SubPath(configPath + "/" + configName)
      // main test.
      ConvertByProfile(mdir.FsSub, conf, profileName, inputListing)
      ExecuteByApplying()
      // remove test file system.
      outputListing := mdir.ListTree()
      mdir.CleanUp()
      // return.
      return outputListing
    },
    Tests:
    []gt.TestList{
      {
        TestName: "main_convert-by-profile_00",
        IsMulti:  true,
        InputArr: []string{
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_ALL,
      },
    },
  },

  /*
  * NOTE: this actually should be a E2E test.
  * Pipe test: ls   | grep -E 'txt$' | renamer -profile files_txt
  * Pipe test: find | grep -E 'txt$' | renamer -profile files_txt
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      firstCommand  := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      profileName   := in.InputArr[3]
      configContent := in.InputArr[4]
      // run test setup.
      path := tu.MakeRealTestFs()
      tu.CreateFile(path, configPath, configName, configContent)
      // simulate pipe.
      cmds := CommandList{
        {
          firstCommand,
          []string{},
        },
        {
          "grep",
          []string{"-E", ".txt$"},
        },
      }
      inputListing := simulatePipe(cmds, path, t)
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      conf          := mdir.SubPath(configPath + "/" + configName)
      // main test.
      _ = ConvertByProfile( // profile command: "renamer -profile 'profileName'"
        mdir.FsSub,         // file system.
        conf,               // path to config.
        profileName,        // profile.
        inputListing,       // input.
      )
      ExecuteByApplying()
      // remove test file system.
      outputListing := mdir.ListTreeOsfs()
      mdir.CleanUp()
      // return.
      return outputListing
    },
    Tests:
    []gt.TestList{
      {
        TestName: "full-test_pipe-test_00",
        IsMulti:  true,
        InputArr: []string{
          "ls",             // first command.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: LS_TXT,
      },
      {
        TestName: "full-test_pipe-test_00",
        IsMulti:  true,
        InputArr: []string{
          "find",           // first command.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_TXT,
      },
    },
  },

  /* Fin test suite. */
}

