
package edit

import (
  // tests.
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
  * Test for editor.
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      editor := NewEditor(
        []func(string) (string, error) {
          func(string) (string, error) { return "1st editing result", nil },
          func(string) (string, error) { return "2nd editing result", nil },
          func(string) (string, error) { return "3rd editing result", nil },
        },
      )
      s, _ := editor.Edit("abc")
      return s
    },
    Tests:
    []gt.TestList{
      {
        TestName: "edit_mock-edits_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue: "1st editing result",
      },
    },
  },

  /* Fin test suite. */
}

