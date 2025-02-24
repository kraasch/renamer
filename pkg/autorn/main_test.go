
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
  * Test for AutoRenamer, Parse(), ConvertWith().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      toml := `title = "My Conf"` + NL +
              `` + NL +
              `[profiles]` + NL +
              `  [profiles.prettify-txt]` + NL +
              `    name = "SomeProfile"` + NL +
              `    [profiles.prettify-txt.profile_rule]` + NL +
              `      word_separators = " ()"` + NL +
              `      delete_chars = ""` + NL +
              `      small_gap_mark = "-"` + NL +
              `      big_gap_mark = "_"` + NL +
              `      conversions = "caA"` + NL +
              `      modes_string = ""` + NL +
              ``
      targetString := in.InputArr[0]
      var auto AutoRenamer
      auto.Parse(toml)
      output := auto.ConvertWith("prettify-txt", targetString, nil)
      return output
    },
    Tests:
    []gt.TestList{
      {
        TestName: "autorn_apply-profile_single-line_file_00",
        IsMulti:  false,
        InputArr: []string{
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_file_00",
        IsMulti:  true,
        InputArr: []string{
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv" + NL +
          "THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_00",
        IsMulti:  true,
        InputArr: []string{
          "videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv" + NL +
          "videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_01",
        IsMulti:  true,
        InputArr: []string{
          "./videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv" + NL +
          "./videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_02",
        IsMulti:  true,
        InputArr: []string{
          "./Videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./Videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./Videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv" + NL +
          "./Videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_03",
        IsMulti:  true,
        InputArr: []string{
          "./my videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./my videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./my videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv" + NL +
          "./my videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_04",
        IsMulti:  true,
        InputArr: []string{
          "./my videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./my Videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./my videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv" + NL +
          "./my Videos/THE-WALKING-DEAD-S05E01-NO-SANCTUARY_1080P-X265-JOY.mkv",
      },
    },
  },

  /*
  * Test for AutoRenamer, Parse(), ConvertWith().
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      toml := `title = "My Conf"` + NL +
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
              ``
      targetString := in.InputArr[0]
      var auto AutoRenamer
      auto.Parse(toml)
      output := auto.ConvertWith("prettify-txt", targetString, nil)
      return output
    },
    Tests:
    []gt.TestList{
      {
        TestName: "autorn_apply-profile_single-line_file_00",
        IsMulti:  false,
        InputArr: []string{
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_file_00",
        IsMulti:  true,
        InputArr: []string{
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv" + NL +
          "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_00",
        IsMulti:  true,
        InputArr: []string{
          "videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv" + NL +
          "videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_01",
        IsMulti:  true,
        InputArr: []string{
          "./videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv" + NL +
          "./videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_02",
        IsMulti:  true,
        InputArr: []string{
          "./Videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./Videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./Videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv" + NL +
          "./Videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_03",
        IsMulti:  true,
        InputArr: []string{
          "./my videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./my videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./my videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv" + NL +
          "./my videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "autorn_apply-profile_multi-line_path_04",
        IsMulti:  true,
        InputArr: []string{
          "./my videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv" + NL +
          "./my Videos/The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
          "./my videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv" + NL +
          "./my Videos/the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
    },
  },

  /* Fin test suite. */
}

