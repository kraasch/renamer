
package e2e

import (
  // unit test.
  "fmt"
  "testing"
  gt "github.com/kraasch/gotest/gotest"
  // e2e test.
  "os/exec"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

func BuildBinary() error {
  fmt.Println("Building.")
  cmd := exec.Command("sh", "-c", "cd .. && make build")
  err := cmd.Run()
  return err
}

var suites = []gt.TestSuite{
  /*
  * Test for the XXX().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      err := BuildBinary()
      if err != nil {
        return "ERROR: Failed to build"
      }
      return ""
    },
    Tests:
    []gt.TestList{
      {
        TestName: "e2e_basic_00",
        IsMulti:  true,
        InputArr: []string{},
        ExpectedValue: "toast",
      },
    },
  },
  /* Fin test suite. */
}

