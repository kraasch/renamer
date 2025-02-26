
package configurator

import (
  // unit testing.
  "testing"
  gt "github.com/kraasch/gotest/gotest"

  // local packages.
  tu "github.com/kraasch/renamer/pkg/testutil"

  // misc packages.
  "fmt"
)

var (
  NL = fmt.Sprintln()
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

/*
* TODO: create tests:
* - [ ] OpenDefaultConfig():
*   - provide with default config.
*   - provide with default config location.
*   - create config if not exists.
*   - read config file text as raw string.
*/

var suites = []gt.TestSuite{

  /*
  * Test configurator xxx.
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      configPath    := in.InputArr[0]
      configName    := in.InputArr[1]
      configContent := in.InputArr[2]
      // run test setup.
      path := tu.MakeEmptyRealTestFs()
      full := path + "/" + configPath
      CreateFile(full, configName, configContent)
      // start test.
      output := ReadConfig(full + "/" + configName)
      // clean up test setup.
      tu.CleanUpRealTestFs(path)
      // return.
      return output
    },
    Tests:
    []gt.TestList{
      {
        TestName: "configurator_basic-test_00",
        IsMulti:  true,
        InputArr: []string{
          "config",         // config path.
          "general.config", // config name.
          // config content.
          `title = "Basic Conf"                     ` + NL +
          `[profiles]                               ` + NL +
          `  [profiles.prettify-txt]                ` + NL +
          `    name = "SomeProfile"                 ` + NL +
          `    [profiles.prettify-txt.profile_rule] ` + NL +
          `      word_separators = " ()"            ` + NL +
          `      delete_chars    = ""               ` + NL +
          `      small_gap_mark  = "-"              ` + NL +
          `      big_gap_mark    = "_"              ` + NL +
          `      conversions     = "caA"            ` + NL +
          `      modes_string    = ""               ` + NL +
          `                                         `,
        },
        ExpectedValue:
          `title = "Basic Conf"                     ` + NL +
          `[profiles]                               ` + NL +
          `  [profiles.prettify-txt]                ` + NL +
          `    name = "SomeProfile"                 ` + NL +
          `    [profiles.prettify-txt.profile_rule] ` + NL +
          `      word_separators = " ()"            ` + NL +
          `      delete_chars    = ""               ` + NL +
          `      small_gap_mark  = "-"              ` + NL +
          `      big_gap_mark    = "_"              ` + NL +
          `      conversions     = "caA"            ` + NL +
          `      modes_string    = ""               ` + NL +
          `                                         `,
      },
    },
  },

  /* Fin test suite. */
}

