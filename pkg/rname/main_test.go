
package rname

import (

  // this is a test.
  "testing"

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
  inputArr2         []string
  expectedValue     string
}

type TestSuite struct {
  testingFunction   func(in TestList) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the function Toast().
  */
  {
    testingFunction:
    func(in TestList) (out string) {
      inputValue := in.inputArr[0]
      out = Toast(inputValue)
      return
    },
    tests:
    []TestList{
      {
        testName:      "category_description_number00",
        isMulti:       false,
        inputArr:      []string{"is nice!"},
        expectedValue: "\"Toast: is nice!\"",
      },
    },
  },
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

