
package fsmanage

import (
  // tests.
  "fmt"
  "testing"
  gt "github.com/kraasch/gotest/gotest"
  "github.com/spf13/afero"

  // local packages.
  dir "github.com/kraasch/renamer/pkg/dir"
  tu "github.com/kraasch/renamer/pkg/testutil"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}


var suites = []gt.TestSuite{
  /*
  * Test for the DirRename().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) (out string) {
      originalNames := in.InputArr[0]
      targetNames   := in.InputArr[1]
      fileSystem := tu.MakeTestFs()
      DirRename(fileSystem, originalNames, targetNames)
      fs2 := afero.NewIOFS(fileSystem)
      out = dir.DirListTree(fs2)
      return
    },
    Tests:
    []gt.TestList{
      {
        TestName: "use-afero_00",
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

