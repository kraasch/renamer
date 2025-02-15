
package dir

import (

  // this is a test.
  "testing"
  "testing/fstest"

  // printing and formatting.
  "fmt"

  // other imports.
  "github.com/kraasch/godiff/godiff"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  isMulti           bool
  inputArr          []string
  expectedValue     string
}

type TestSuite struct {
  testingFunction   func(in TestList) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the DirList().
  */
  {
    testingFunction:
    func(in TestList) (out string) {
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
    tests:
    []TestList{
      {
        testName: "dir_test-01_list-tree_00",
        isMulti:  true,
        inputArr: []string{},
        expectedValue:
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
    testingFunction:
    func(in TestList) (out string) {
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
    tests:
    []TestList{
      {
        testName: "dir_test-01_list-tree_00",
        isMulti:  true,
        inputArr: []string{},
        expectedValue:
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

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      t.Run(name, func(t *testing.T) {
        exp := test.expectedValue
        got := suite.testingFunction(test)
        if exp != got {
          if test.isMulti {
            t.Errorf("In '%s':\n", name)
            diff := godiff.CDiff(exp, got)
            t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
            t.Errorf("exp/got:\n%s\n", diff)
          } else {
            t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
          }
        }
      })
    }
  }
}

