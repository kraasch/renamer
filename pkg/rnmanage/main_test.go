
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
  for i, c := range commands {
    cmd := exec.Command(c.Name, c.Args...)
    cmd.Dir = path // execute within diretory of test file system.
    cmd.Stdin = strings.NewReader(string(output))
    output, _ = cmd.Output()
    fmt.Printf("%d > %s %v \t==> %s\n", i, c.Name, c.Args, output)
  }
  return string(output)
}

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
      configPath    := in.InputArr[0]
      configName    := in.InputArr[1]
      profileName   := in.InputArr[2]
      configContent := in.InputArr[3]

      // run test setup.
      path := tu.MakeRealTestFs()
      tu.CreateFile(path, configPath, configName, configContent)

      // simulate pipe.
      cmds := CommandList{
        {
          "ls",
          []string{},
        },
        {
          "grep",
          []string{"-E", ".txt$"},
        },
      }
      finalPipeOutput := simulatePipe(cmds, path)

      output := Command(
        "testfs/" + configPath + "/" + configName, // config.
        "profile",       // type (profile/edit/interactive).
        profileName,     // profile.
        finalPipeOutput, // input.
      )

      // clean up test setup.
      tu.CleanUpRealTestFs(path)

      // return.
      return output
    },
    Tests:
    []gt.TestList{
      {
        TestName: "full-test_pipe-test_00",
        IsMulti:  true,
        InputArr: []string{
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          // config content.
          `# Basic config              ` + NL +
          `                            ` + NL +
          `title = "TOML Example"      ` + NL +
          `                            ` + NL +
          `[profiles]                  ` + NL +
          `                            ` + NL +
          `    [profiles.toast-txt]    ` + NL +
          `    name = "toast-txt"      ` + NL +
          `    rule = "AAA"            ` + NL +
          `                            ` + NL +
          `    [profiles.prettify-txt] ` + NL +
          `    name = "prettify-txt"   ` + NL +
          `    rule = "more pipe tests"` + NL +
          `                            `,
        },
        ExpectedValue: "Implement", // TODO: continue here.
        // "fruits/"             + NL +
        // "fruits/APPLES.txt"   + NL +
        // "fruits/BANANAS.txt"  + NL +
        // "fruits/COCONUTS.txt" + NL +
        // "NOTES.txt"           + NL +
        // "shapes/"             + NL +
        // "shapes/CIRCLE.txt"   + NL +
        // "shapes/SQUARE.txt"   + NL +
        // "shapes/TRIANGLE.txt",
      },
    },
  },
  /* Fin test suite. */
}

