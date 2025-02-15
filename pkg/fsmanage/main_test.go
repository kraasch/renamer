
package fsmanage

import (
  // tests.
  "fmt"
  "testing"
  "testing/fstest"
  gt "github.com/kraasch/gotest/gotest"

  // local packages.
  dir "github.com/kraasch/renamer/pkg/dir"
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
    func(in gt.TestList) (out string) {
      targetNames := in.InputArr[0]
      testFs := fstest.MapFS{
        "notes.txt":           {Data: []byte("")},
        "fruits/apples.txt":   {Data: []byte("")},
        "fruits/bananas.txt":  {Data: []byte("")},
        "shapes/triangle.txt": {Data: []byte("")},
        "fruits/coconuts.txt": {Data: []byte("")},
        "shapes/square.txt":   {Data: []byte("")},
        "shapes/circle.txt":   {Data: []byte("")},
      }
      DirRename(testFs, targetNames)
      out = dir.DirListTree(testFs)
      return
    },
    Tests:
    []gt.TestList{
      {
        TestName: "dir_list-tree_00",
        IsMulti:  true,
        InputArr: []string{
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

