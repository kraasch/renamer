
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
  * Test for the ReadRawProfileConfig().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      configText := in.InputArr[0]
      config := ReadRawProfileConfig(configText)
      profile := config.Profiles["prettify-txt"]
      return profile.Rule
    },
    Tests:
    []gt.TestList{
      {
        TestName: "profiler_read-profiles_from-config_00",
        IsMulti:  true,
        InputArr: []string{
          "# My config" + NL +
          "" + NL +
          "title = \"TOML Example\"" + NL +
          "" + NL +
          "[profiles]" + NL +
          "" + NL +
          "    [profiles.toast-txt]" + NL +
          "    name = \"toast-txt\"" + NL +
          "    rule = \"ZZZ\"" + NL +
          "" + NL +
          "    [profiles.prettify-txt]" + NL +
          "    name = \"prettify-txt\"" + NL +
          "    rule = \"YYY\"" + NL +
          "",
        },
        ExpectedValue: "YYY",
      },
    },
  },
  /* Fin test suite. */
}

