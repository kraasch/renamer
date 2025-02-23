
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
  * Test for xxx.
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      toml := in.InputArr[0]
      name := in.InputArr[1]
      var auto AutoRenamer
      auto.Parse(toml)
      output := auto.ConvertWith("prettify-txt", name)
      return output
    },
    Tests:
    []gt.TestList{
      {
        TestName: "autorn_apply-profile_00",
        IsMulti:  false,
        InputArr: []string{
          `title = "My Conf"` + NL +
          `` + NL +
          `[profiles]` + NL +
          `  [profiles.prettify-txt]` + NL +
          `    name = "SomeProfile"` + NL +
          `    [profiles.prettify-txt.profile_rule]` + NL +
          `      word_separators = " ()"` + NL +
          `      delete_chars = ""` + NL +
          `      small_gap_mark = "-"` + NL +
          `      big_gap_mark = "_"` + NL +
          `      conversions = "cAa"` + NL +
          `      modes_string = ""` + NL +
          ``,
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
    },
  },

  /* Fin test suite. */
}

