
package autorn

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
  * Test for the xxx().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      return "the auto-renamer"
    },
    Tests:
    []gt.TestList{
      {
        TestName: "auto-renamer_stub_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue:
              "Implement",
      },
    },
  },
  /* Fin test suite. */
}

