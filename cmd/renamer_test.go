
package main

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
  rn "github.com/kraasch/renamer/pkg/rnmanage"
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
      profileName   := in.InputArr[1]
      configContent := in.InputArr[2]

      // run test setup.
      path := tu.MakeRealTestFs()

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

      cmd := rn.Command(
        configPath,      // config.
        "profile",       // type (profile/edit/interactive).
        profileName,     // profile.
        finalPipeOutput, // input.
      )
      // output := cmd.Output()
      output := cmd

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
          ".config/renamer/general.config", // config path.
          "prettify-txt", // profile name.
          // config content.
          "# My config" + NL +
          "" + NL +
          "title = \"TOML Example\"" + NL +
          "" + NL +
          "[owner]" + NL +
          "name = \"Tom Preston-Werner\"" + NL +
          "dob = 1979-05-27T07:32:00-08:00" + NL +
          "" + NL +
          "[database]" + NL +
          "enabled = true" + NL +
          "ports = [ 8000, 8001, 8002 ]" + NL +
          "data = [ [\"delta\", \"phi\"], [3.14] ]" + NL +
          "temp_targets = { cpu = 79.5, case = 72.0 }",
        },
        ExpectedValue:
        "FRUITS/" + NL +
        "FRUITS/apples.txt" + NL +
        "FRUITS/bananas.txt" + NL +
        "FRUITS/coconuts.txt" + NL +
        "NOTES.txt" + NL +
        "Shapes/" + NL +
        "Shapes/circle.txt" + NL +
        "Shapes/square.txt" + NL +
        "Shapes/triangle.txt",
      },
    },
  },
  /* Fin test suite. */
}

