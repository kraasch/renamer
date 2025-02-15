
package dir

import (
  "fmt"
  "testing"
  "testing/fstest"
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
  * Test for the DirList().
  */
  {
    TestingFunction:
    func(in gt.TestList) (out string) {
      testFs := fstest.MapFS{
        "notes.txt":           {Data: []byte("")},
        "fruits/apples.txt":   {Data: []byte("")},
        "fruits/bananas.txt":  {Data: []byte("")},
        "shapes/triangle.txt": {Data: []byte("")},
        "fruits/coconuts.txt": {Data: []byte("")},
        "shapes/square.txt":   {Data: []byte("")},
        "shapes/circle.txt":   {Data: []byte("")},
      }
      out = DirList(testFs)
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
    func(in gt.TestList) (out string) {
      testFs := fstest.MapFS{
        "notes.txt":           {Data: []byte("")},
        "fruits/apples.txt":   {Data: []byte("")},
        "fruits/bananas.txt":  {Data: []byte("")},
        "shapes/triangle.txt": {Data: []byte("")},
        "fruits/coconuts.txt": {Data: []byte("")},
        "shapes/square.txt":   {Data: []byte("")},
        "shapes/circle.txt":   {Data: []byte("")},
      }
      out = DirListTree(testFs)
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

