
package profiler

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
  * Test for the DirList().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      return "Profiler"
    },
    Tests:
    []gt.TestList{
      {
        TestName: "edit_stub_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue:
              "Profiler!",
      },
    },
  },
  /* Fin test suite. */
}

