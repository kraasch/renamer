
package rnmanage

import (
  "testing"
  gt "github.com/kraasch/gotest/gotest"
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

var suites = []gt.TestSuite{
  /*
  * Test for the toast().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      name := in.InputArr[0]
      out := Toast(name)
      return out
    },
    Tests:
    []gt.TestList{
      {
        TestName: "rn_test-stub_00",
        IsMulti:  true,
        InputArr: []string{ "Joe" },
        ExpectedValue: "Smello Joe?",
      },
    },
  },
  /* Fin test suite. */
}

