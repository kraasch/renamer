
package autorn

import (
  // tests.
  "fmt"
  "testing"
  gt "github.com/kraasch/gotest/gotest"

  // local packages.
  // pro "github.com/kraasch/renamer/pkg/profiler"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

var suites = []gt.TestSuite{

  /*
  * Test for xxx.
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // wordSeparators := in.InputArr[0]
      // deleteChars    := in.InputArr[1]
      // smallGapMark   := in.InputArr[2]
      // bigGapMark     := in.InputArr[3]
      // conversions    := in.InputArr[4]
      // modesString    := in.InputArr[5]
      // targetName     := in.InputArr[6]
      // p := pro.CreateProfile
      output := "Toast"
      return output
    },
    Tests:
    []gt.TestList{
      {
        TestName: "xxx_00",
        IsMulti:  false,
        InputArr: []string{
          " ()", // word separators.
          "",    // delete characters.
          "-",   // small gap replacement.
          "_",   // big gap replacement.
          "cAa", // list of actions.
          "",    // string of modes.
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
          "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
        },
        ExpectedValue: "false",
      },
    },
  },

  /* Fin test suite. */
}

