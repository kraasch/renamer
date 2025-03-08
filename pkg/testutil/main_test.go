
package testutil

import (
  // for tests.
  "fmt"
  "testing"
  gt "github.com/kraasch/gotest/gotest"
  // misc.
  "strings"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

var suites = []gt.TestSuite{

  /*
  * Test for XXX.
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      input := in.InputArr[0]
      return strings.ToUpper(input)
    },
    Tests:
    []gt.TestList{
      {
        TestName: "test-util_00",
        IsMulti:  true,
        InputArr: []string{
          "b",
        },
        ExpectedValue:
          "b",
      },
    },
  },

  /* Fin test suite. */
}

