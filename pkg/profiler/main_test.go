
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
      profileName := in.InputArr[0]
      configText  := in.InputArr[1]
      config      := ReadRawProfileConfig(configText)
      profile     := config.Profiles[profileName]
      rule        := profile.ProfileRule
      return rule.WordSeparators
    },
    Tests:
    []gt.TestList{
      {
        TestName: "profiler_read-profiles_from-config_00",
        IsMulti:  true,
        InputArr: []string{
          "prettify-txt",
          `# Some comment                           ` + NL +
          `                                         ` + NL +
          `title = "Some Example"                   ` + NL +
          `                                         ` + NL +
          `[profiles]                               ` + NL +
          `[profiles.toast-txt]                     ` + NL +
          `    name            = "n0"               ` + NL +
          `    [profiles.toast-txt.profile_rule]    ` + NL +
          `    word_separators = "A"                ` + NL +
          `    delete_chars    = "B"                ` + NL +
          `    small_gap_mark  = "C"                ` + NL +
          `    big_gap_mark    = "D"                ` + NL +
          `    conversions     = "E"                ` + NL +
          `    modes_string    = "F"                ` + NL +
          `[profiles.prettify-txt]                  ` + NL +
          `    name            = "n1"               ` + NL +
          `    [profiles.prettify-txt.profile_rule] ` + NL +
          `    word_separators = "YYY"              ` + NL +
          `    delete_chars    = "b"                ` + NL +
          `    small_gap_mark  = "c"                ` + NL +
          `    big_gap_mark    = "d"                ` + NL +
          `    conversions     = "e"                ` + NL +
          `    modes_string    = "f"                ` + NL +
          `                                         `,
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
      r := Rule{
        WordSeparators: in.InputArr[0],
        DeleteChars:    in.InputArr[1],
        SmallGapMark:   in.InputArr[2],
        BigGapMark:     in.InputArr[3],
        Conversions:    in.InputArr[4],
        ModesString:    in.InputArr[5],
      }
      profileName    := in.InputArr[6]
      testName       := in.InputArr[7]
      p := Profile{ profileName, r }
      output := p.Apply(testName)
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
          "SomeProfile", // profile name.
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
        "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
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

