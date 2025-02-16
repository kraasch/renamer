
package dir

import (
  // tests.
  "fmt"
  "testing"
  gt "github.com/kraasch/gotest/gotest"
  "github.com/spf13/afero"

  // local packages.
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
  * Test for the DirList().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) (out string) {
      fs1 := tu.MakeTestFs()
      fs2 := afero.NewIOFS(fs1)
      out = DirList(fs2)
      return
    },
    Tests:
    []gt.TestList{
      {
        TestName: "dir_list-dir_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue:
              "fruits/" + NL +
              "notes.txt" + NL +
              "shapes/",
      },
    },
  },
  /*
  * Test for the DirListTree().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) (out string) {
      fs1 := tu.MakeTestFs()
      fs2 := afero.NewIOFS(fs1)
      out = DirListTree(fs2)
      return
    },
    Tests:
    []gt.TestList{
      {
        TestName: "dir_list-tree_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue:
              "fruits/" + NL +
              "fruits/apples.txt" + NL +
              "fruits/bananas.txt" + NL +
              "fruits/coconuts.txt" + NL +
              "notes.txt" + NL +
              "shapes/" + NL +
              "shapes/circle.txt" + NL +
              "shapes/square.txt" + NL +
              "shapes/triangle.txt",
      },
    },
  },
  /* Fin test suite. */
}

