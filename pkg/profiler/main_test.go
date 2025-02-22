
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

  /*
  * Test for CreateProfile().
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

  /*
  * Test for AddProfileToConfig().
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

