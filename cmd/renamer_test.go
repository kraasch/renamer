
package main

import (
  "fmt"
  "testing"
  gt "github.com/kraasch/gotest/gotest"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

var suites = []gt.TestSuite{
  /*
  * Test for the XXXXXXXXXXXX().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // originalNames := in.InputArr[0]
      // targetNames   := in.InputArr[1]
      // fileSystem := tu.MakeTestFs()
      // XXXXXXXXXXXX(fileSystem, originalNames, targetNames)
      // fs2 := afero.NewIOFS(fileSystem)
      // out = dir.DirListTree(fs2)
      // return out
      return "XXXXXXXXXXXX"
    },
    Tests:
    []gt.TestList{
      {
        TestName: "full-test_pipe-idea_00",
        IsMulti:  true,
        InputArr: []string{
          "fruits/" + NL +
          "notes.txt" + NL +
          "shapes/",
          "FRUITS/" + NL +
          "NOTES.txt" + NL +
          "Shapes/",
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

