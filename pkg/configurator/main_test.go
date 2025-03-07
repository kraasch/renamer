
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
  BASIC_CONF =
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
    `                                         `
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
  * Test AutoReadConfig().
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
      // start test.
      c := Configurator{
        ConfigFileName: configName,
        PathToConfig:   configPath,
        DefaultConfig:  configContent,
      }
      c.SetRoot(path)
      output := c.AutoReadConfig()
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
          ".config/renamer/", // config path.
          "renamer.toml",     // config name.
          BASIC_CONF,         // config content.
        },
        ExpectedValue:
          BASIC_CONF,
      },
    },
  },

  /*
  * Test configurator:
  *  - Configurator{}
  *  - SetRoot()
  *  - ExistsDefaultConfig()
  *  - CreateDefaultConfig()
  *  - ReadConfig()
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
      // start test.
      c := Configurator{
        ConfigFileName: configName,
        PathToConfig:   configPath,
        DefaultConfig:  configContent,
      }
      c.SetRoot(path)
      if !c.ExistsDefaultConfig() {
        c.CreateDefaultConfig()
      }
      output := c.ReadConfig()
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
          ".config/renamer/", // config path.
          "renamer.toml",     // config name.
          BASIC_CONF,         // config content.
        },
        ExpectedValue:
          BASIC_CONF,
      },
    },
  },

  /*
  * Test configurator CreateFile(), ReadConfig().
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
      // start test.
      CreateFile(full, configName, configContent)
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
          BASIC_CONF,       // config content.
        },
        ExpectedValue:
          BASIC_CONF,
      },
    },
  },

  /* Fin test suite. */
}

