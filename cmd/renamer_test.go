
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

      profile := rn.Profile( // TODO: create profile here.
        in.InputArr[0], // PROFILE DATA 0
        in.InputArr[1], // PROFILE DATA 1
      )
      cmd := rn.Command(
        finalPipeOutput, // stdin.
        "profile",       // type.
        profile,         // profile.
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
          "PROFILE DATA 0",
          "PROFILE DATA 1",
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

