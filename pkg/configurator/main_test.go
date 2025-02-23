
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
      profileName   := in.InputArr[2]
      configContent := in.InputArr[3]
      // run test setup.
      path := tu.MakeRealTestFs()
      // TODO: make a copy of CreateFile():
      // - make it part of the configurator package.
      // - make it take a path to config, not only single folder.
      tu.CreateFile(path, configPath, configName, configContent)
      // start test.
      output := Toast(configPath + configName + profileName)
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
          "prettify-txt",   // profile name.
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
          "./fruits/"             + NL +
          "./fruits/APPLES.txt"   + NL +
          "./fruits/BANANAS.txt"  + NL +
          "./fruits/COCONUTS.txt" + NL +
          "./NOTES.txt"           + NL +
          "./shapes/"             + NL +
          "./shapes/CIRCLE.txt"   + NL +
          "./shapes/SQUARE.txt"   + NL +
          "./shapes/TRIANGLE.txt" + NL,
      },
    },
  },

  /* Fin test suite. */
}

