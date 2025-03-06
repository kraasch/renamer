
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

/*
* TODO: use exec.Command(name, args...) with
* vim in order to do automated tests with mocks.
*   - "vim --headless +qa"
*   - "vi -c 'normal gg0llllrxu' some-file.txt || (echo 'Error!')"
* Tasks:
*   - [ ] test for successful edits.
*   - [ ] test for empty edits.
*   - [ ] test for aborted edits.
*/

var suites = []gt.TestSuite{

  /*
  * Test for editor.
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // mEditor := NewMockEditor(
      //   []func(string) (string, error) {
      //     func(string) (string, error) { return "1st editing result", nil },
      //     func(string) (string, error) { return "2nd editing result", nil },
      //     func(string) (string, error) { return "3rd editing result", nil },
      //   },
      // )
      // s, _ := mEditor.editor.Edit("abc")
      // return s
      return "NOT IMPLEMENTNED YET"
    },
    Tests:
    []gt.TestList{
      {
        TestName: "edit_mock-edits_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue: "NOT IMPLEMENTNED YET",
      },
    },
  },

  /* Fin test suite. */
}

