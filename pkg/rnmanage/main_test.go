
package rnmanage

import (
  // unit testing.
  "testing"
  gt "github.com/kraasch/gotest/gotest"
  "github.com/spf13/afero"

  // cli testing.
  "fmt"
  "os/exec"
  "strings"

  // misc.
  "path/filepath"
  "os"

  // local packages.
  tu "github.com/kraasch/renamer/pkg/testutil"
  dir "github.com/kraasch/renamer/pkg/dir"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

type CommandList []struct {
  Name string
  Args []string
}

func simulatePipe(commands CommandList, path string) string {
  var output []byte
  // for i, c := range commands { // TODO: for logging.
  for _, c := range commands {
    cmd := exec.Command(c.Name, c.Args...)
    cmd.Dir = path // execute within diretory of test file system.
    cmd.Stdin = strings.NewReader(string(output))
    output, _ = cmd.Output()
    // TODO: make this into a log or something.
    // fmt.Printf("%d > %s %v \t==> %s\n", i, c.Name, c.Args, output)
  }
  return string(output)
}

/*
* TODO: do not make too many pipe tests, but eventually implement some e2e tests.
*/

/*
* NOTE: PIPE TEST IDEAS:
* Pipe test: ls | grep -E 'mp3$' | renamer -edit // TODO: test.
* Pipe test: find | grep -E '.ogg$' | renamer -profile music_ogg // TODO: test.
*/

var suites = []gt.TestSuite{

  /*
  * Pipe test: ls | grep -E 'txt$' | renamer -profile files_txt
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
      finalPipeOutput := simulatePipe(cmds, path)
      // create test file system.

      fs := afero.NewOsFs()
      currentDir, _ := os.Getwd()
      targetDir := filepath.Join(currentDir, "testfs")
      fs4 := afero.NewBasePathFs(fs, targetDir)

      // main test.
      // TODO: implement command types: profile, edit, interactive.
      _ = Command(
        fs4, // file system.
        "testfs/" + configPath + "/" + configName, // config.
        profileName,     // profile.
        finalPipeOutput, // input.
      )
      // get file listing.
      fs2 := afero.NewIOFS(fs)
      fs3,_ := fs2.Sub("testfs")
      listing := dir.DirListTree(fs3)
      // listing := tu.ListFs(fs, "testfs/") // TODO: remove later.
      // clean up test setup.
      tu.CleanUpRealTestFs(path)
      // return.
      return listing
    },
    Tests:
    []gt.TestList{
      {
        TestName: "full-test_pipe-test_00",
        IsMulti:  true,
        InputArr: []string{
          "find",           // first command.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          // config content.
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
          `                                         `,
        },
        ExpectedValue:
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
        "shapes/TRIANGLE.txt"   + NL,
      },
    },
  },

  /* Fin test suite. */
}

