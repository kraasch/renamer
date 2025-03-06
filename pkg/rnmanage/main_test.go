
package rnmanage

import (
  // unit testing.
  "testing"
  gt "github.com/kraasch/gotest/gotest"

  // cli testing.
  "fmt"
  "os/exec"
  "strings"

  // local packages.
  tu "github.com/kraasch/renamer/pkg/testutil"
)

var (
  NL = fmt.Sprintln()
  BASIC_CONF =
  `title = "Basic Conf"                     ` + NL +
  `[profiles]                               ` + NL +
  `  [profiles.lcase-txt]                   ` + NL +
  `    name = "LowerCase"                   ` + NL +
  `    [profiles.lcase-txt.profile_rule]    ` + NL +
  `      word_separators = " ()"            ` + NL +
  `      delete_chars    = ""               ` + NL +
  `      small_gap_mark  = "-"              ` + NL +
  `      big_gap_mark    = "_"              ` + NL +
  `      conversions     = "cAa"            ` + NL +
  `      modes_string    = ""               ` + NL +
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
  FIND_TXT_BANANAS =
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/apples.txt"     + NL +
  "fruits/coconuts.txt"   + NL +
  "notes.txt"             + NL +
  "shapes/"               + NL +
  "shapes/circle.txt"     + NL +
  "shapes/square.txt"     + NL +
  "shapes/triangle.txt"
  FIND_TXT_BANANAS_2 =
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/apples.txt"     + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/coconuts.txt"   + NL +
  "notes.txt"             + NL +
  "shapes/"               + NL +
  "shapes/circle.txt"     + NL +
  "shapes/square.txt"     + NL +
  "shapes/triangle.txt"
  FIND_TXT =
  "NOTES.txt"             + NL +
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/APPLES.txt"     + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/COCONUTS.txt"   + NL +
  "shapes/"               + NL +
  "shapes/CIRCLE.txt"     + NL +
  "shapes/SQUARE.txt"     + NL +
  "shapes/TRIANGLE.txt"
  FIND_TXT_FRUITS_ONLY =
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/APPLES.txt"     + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/COCONUTS.txt"   + NL +
  "notes.txt"             + NL +
  "shapes/"               + NL +
  "shapes/circle.txt"     + NL +
  "shapes/square.txt"     + NL +
  "shapes/triangle.txt"
  FIND_TXT_SHAPES_ONLY =
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/apples.txt"     + NL +
  "fruits/bananas.txt"    + NL +
  "fruits/coconuts.txt"   + NL +
  "notes.txt"             + NL +
  "shapes/"               + NL +
  "shapes/CIRCLE.txt"     + NL +
  "shapes/SQUARE.txt"     + NL +
  "shapes/TRIANGLE.txt"
  LS_TXT =
  "NOTES.txt"             + NL +
  "config/"               + NL +
  "config/general.config" + NL +
  "fruits/"               + NL +
  "fruits/apples.txt"     + NL +
  "fruits/bananas.txt"    + NL +
  "fruits/coconuts.txt"   + NL +
  "shapes/"               + NL +
  "shapes/circle.txt"     + NL +
  "shapes/square.txt"     + NL +
  "shapes/triangle.txt"
  FIND_ALL =
  "NOTES.txt"             + NL +
  "config/"               + NL +
  "config/GENERAL.config" + NL +
  "fruits/"               + NL +
  "fruits/APPLES.txt"     + NL +
  "fruits/BANANAS.txt"    + NL +
  "fruits/COCONUTS.txt"   + NL +
  "shapes/"               + NL +
  "shapes/CIRCLE.txt"     + NL +
  "shapes/SQUARE.txt"     + NL +
  "shapes/TRIANGLE.txt"
  FIND_ALL_FORMATTED =
  "config/               => config/"               + NL +
  "config/general.config => config/GENERAL.config" + NL +
  "fruits/               => fruits/"               + NL +
  "fruits/apples.txt     => fruits/APPLES.txt"     + NL +
  "fruits/bananas.txt    => fruits/BANANAS.txt"    + NL +
  "fruits/coconuts.txt   => fruits/COCONUTS.txt"   + NL +
  "notes.txt             => NOTES.txt"             + NL +
  "shapes/               => shapes/"               + NL +
  "shapes/circle.txt     => shapes/CIRCLE.txt"     + NL +
  "shapes/square.txt     => shapes/SQUARE.txt"     + NL +
  "shapes/triangle.txt   => shapes/TRIANGLE.txt"   + NL
  FIND_BANANA_FORMATTED =
  "notes.txt             => notes.txt"             + NL +
  "config/               => config/"               + NL +
  "config/general.config => config/general.config" + NL +
  "fruits/               => fruits/"               + NL +
  "fruits/apples.txt     => fruits/apples.txt"     + NL +
  "fruits/bananas.txt    => fruits/BANANAS.txt"    + NL +
  "fruits/coconuts.txt   => fruits/coconuts.txt"   + NL +
  "shapes/               => shapes/"               + NL +
  "shapes/circle.txt     => shapes/circle.txt"     + NL +
  "shapes/square.txt     => shapes/square.txt"     + NL +
  "shapes/triangle.txt   => shapes/triangle.txt"
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

type CommandList []struct {
  Name string
  Args []string
}

func simulatePipe(commands CommandList, path string, t *testing.T) string {
  var output []byte
  for i, c := range commands {
    cmd := exec.Command(c.Name, c.Args...)
    cmd.Dir = path // execute within diretory of test file system.
    cmd.Stdin = strings.NewReader(string(output))
    output, _ = cmd.Output()
    t.Logf("%d > %s %v \t==> %s\n", i, c.Name, c.Args, output)
  }
  return string(output)
}

var suites = []gt.TestSuite{

  /*
  * Test ConvertByPathList() and ExecuteByFormatting()
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      workingDir    := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      configContent := in.InputArr[4]
      pathList      := in.InputArr[5]
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      inputListing  := mdir.ListTree()
      // main test.
      ConvertByPathList(
        mdir.FsSub,
        workingDir,
        pathList,
        inputListing,
      )
      formattedOutput := ExecuteByFormatting()
      // remove test file system.
      mdir.CleanUp()
      // return.
      return formattedOutput
    },
    Tests:
    []gt.TestList{
      {
        TestName: "path-list_00",
        IsMulti:  true,
        InputArr: []string{
          ".",                // path to working directory.
          "config",           // config path     // TODO: REMOVE.
          "general.config",   // config name.    // TODO: REMOVE.
          "prettify-txt",     // profile name.   // TODO: REMOVE.
          BASIC_CONF,         // config content. // TODO: REMOVE.
          FIND_TXT_BANANAS_2, // path list.
        },
        ExpectedValue: FIND_BANANA_FORMATTED,
      },
    },
  },

  /*
  * Test ConvertByRule() and ExecuteByFormatting()
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      workingDir    := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      configContent := in.InputArr[4]
      ruleString    := in.InputArr[5]
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      inputListing  := mdir.ListTree()
      // main test.
      ConvertByRule(
        mdir.FsSub,
        workingDir,
        ruleString,
        inputListing,
      )
      formattedOutput := ExecuteByFormatting()
      // remove test file system.
      mdir.CleanUp()
      // return.
      return formattedOutput
    },
    Tests:
    []gt.TestList{
      {
        TestName: "rulestr_00",
        IsMulti:  true,
        InputArr: []string{
          ".",              // path to working directory.
          "config",         // config path     // TODO: REMOVE.
          "general.config", // config name.    // TODO: REMOVE.
          "prettify-txt",   // profile name.   // TODO: REMOVE.
          BASIC_CONF,       // config content. // TODO: REMOVE.
          " ()||-|_|caA|",  // rule string.
        },
        ExpectedValue: FIND_ALL_FORMATTED,
      },
    },
  },

  /*
  * Test ConvertByProfile() and ExecuteByFormatting()
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      workingDir    := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      profileName   := in.InputArr[3]
      configContent := in.InputArr[4]
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      inputListing  := mdir.ListTree()
      conf          := mdir.SubPath(configPath + "/" + configName)
      // main test.
      ConvertByProfile(
        mdir.FsSub,
        workingDir,
        conf,
        profileName,
        inputListing,
      )
      formattedOutput := ExecuteByFormatting()
      // remove test file system.
      mdir.CleanUp()
      // return.
      return formattedOutput
    },
    Tests:
    []gt.TestList{
      {
        TestName: "format_00",
        IsMulti:  true,
        InputArr: []string{
          ".",              // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_ALL_FORMATTED,
      },
    },
  },

  /*
  * Test ConvertByProfile() and ExecuteByValidating()
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      workingDir    := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      profileName   := in.InputArr[3]
      configContent := in.InputArr[4]
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      inputListing  := mdir.ListTree()
      conf          := mdir.SubPath(configPath + "/" + configName)
      // main test.
      ConvertByProfile(
        mdir.FsSub,
        workingDir,
        conf,
        profileName,
        inputListing,
      )
      valid := ExecuteByValidating()
      // remove test file system.
      mdir.CleanUp()
      // return.
      if valid {
        return "true"
      } else {
        return "false"
      }
    },
    Tests:
    []gt.TestList{
      {
        TestName: "validate_00",
        IsMulti:  true,
        InputArr: []string{
          ".",              // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_ALL,
      },
      {
        TestName: "validate_01",
        IsMulti:  true,
        InputArr: []string{
          "./fruits",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_TXT_FRUITS_ONLY,
      },
      {
        TestName: "validate_02",
        IsMulti:  true,
        InputArr: []string{
          "./shapes",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_TXT_SHAPES_ONLY,
      },
      {
        TestName: "validate_no-wdir-dot_00",
        IsMulti:  true,
        InputArr: []string{
          "",               // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_ALL,
      },
      {
        TestName: "validate_no-wdir-dot_01",
        IsMulti:  true,
        InputArr: []string{
          "fruits",         // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_TXT_FRUITS_ONLY,
      },
      {
        TestName: "validate_no-wdir-dot_02",
        IsMulti:  true,
        InputArr: []string{
          "shapes",         // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_TXT_SHAPES_ONLY,
      },
      {
        TestName: "validate_confusing-root-path_00",
        IsMulti:  true,
        InputArr: []string{
          "./fruits/..",    // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_ALL,
      },
      {
        TestName: "validate_confusing-root-path_01",
        IsMulti:  true,
        InputArr: []string{
          "./ABC/..",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "false", //FIND_ALL,
      },
      {
        TestName: "validate_wdir-not-a-dir_00",
        IsMulti:  true,
        InputArr: []string{
          "fruits/bananas.txt", // path to working directory.
          "config",             // config path.
          "general.config",     // config name.
          "prettify-txt",       // profile name.
          BASIC_CONF,           // config content.
        },
        ExpectedValue: "false", //FIND_TXT_BANANAS,
      },
      {
        TestName: "validate_wdir-not-a-dir_01",
        IsMulti:  true,
        InputArr: []string{
          "fruits/bananas",     // path to working directory.
          "config",             // config path.
          "general.config",     // config name.
          "prettify-txt",       // profile name.
          BASIC_CONF,           // config content.
        },
        ExpectedValue: "false", //FIND_TXT_BANANAS,
      },
      {
        TestName: "validate_lower-case_00",
        IsMulti:  true,
        InputArr: []string{
          ".",              // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_ALL,
      },
      {
        TestName: "validate_lower-case_01",
        IsMulti:  true,
        InputArr: []string{
          "./fruits",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_TXT_FRUITS_ONLY,
      },
      {
        TestName: "validate_lower-case_02",
        IsMulti:  true,
        InputArr: []string{
          "./shapes",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_TXT_SHAPES_ONLY,
      },
      {
        TestName: "validate_lower-case_no-wdir-dot_00",
        IsMulti:  true,
        InputArr: []string{
          "",               // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_ALL,
      },
      {
        TestName: "validate_lower-case_no-wdir-dot_01",
        IsMulti:  true,
        InputArr: []string{
          "fruits",         // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_TXT_FRUITS_ONLY,
      },
      {
        TestName: "validate_lower-case_no-wdir-dot_02",
        IsMulti:  true,
        InputArr: []string{
          "shapes",         // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_TXT_SHAPES_ONLY,
      },
      {
        TestName: "validate_lower-case_confusing-root-path_00",
        IsMulti:  true,
        InputArr: []string{
          "./fruits/..",    // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_ALL,
      },
      {
        TestName: "validate_lower-case_confusing-root-path_01",
        IsMulti:  true,
        InputArr: []string{
          "./ABC/..",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "lcase-txt",      // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: "true", //FIND_ALL,
      },
      {
        TestName: "validate_lower-case_wdir-not-a-dir_00",
        IsMulti:  true,
        InputArr: []string{
          "fruits/bananas.txt", // path to working directory.
          "config",             // config path.
          "general.config",     // config name.
          "lcase-txt",          // profile name.
          BASIC_CONF,           // config content.
        },
        ExpectedValue: "true", //FIND_TXT_BANANAS,
      },
      {
        TestName: "validate_lower-case_wdir-not-a-dir_01",
        IsMulti:  true,
        InputArr: []string{
          "fruits/bananas",     // path to working directory.
          "config",             // config path.
          "general.config",     // config name.
          "lcase-txt",          // profile name.
          BASIC_CONF,           // config content.
        },
        ExpectedValue: "true", //FIND_TXT_BANANAS,
      },
    },
  },

  /*
  * Test ConvertByProfile() and ExecuteByApplying()
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      workingDir    := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      profileName   := in.InputArr[3]
      configContent := in.InputArr[4]
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      inputListing  := mdir.ListTree()
      conf          := mdir.SubPath(configPath + "/" + configName)
      // main test.
      ConvertByProfile(
        mdir.FsSub,
        workingDir,
        conf,
        profileName,
        inputListing,
      )
      ExecuteByApplying()
      // remove test file system.
      outputListing := mdir.ListTree()
      mdir.CleanUp()
      // return.
      return outputListing
    },
    Tests:
    []gt.TestList{
      {
        TestName: "main_convert-by-profile_00",
        IsMulti:  true,
        InputArr: []string{
          ".",              // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_ALL,
      },
      {
        TestName: "main_convert-by-profile_01",
        IsMulti:  true,
        InputArr: []string{
          "./fruits",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_TXT_FRUITS_ONLY,
      },
      {
        TestName: "main_convert-by-profile_02",
        IsMulti:  true,
        InputArr: []string{
          "./shapes",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_TXT_SHAPES_ONLY,
      },
      {
        TestName: "main_convert-by-profile_no-wdir-dot_00",
        IsMulti:  true,
        InputArr: []string{
          "",               // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_ALL,
      },
      {
        TestName: "main_convert-by-profile_no-wdir-dot_01",
        IsMulti:  true,
        InputArr: []string{
          "fruits",         // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_TXT_FRUITS_ONLY,
      },
      {
        TestName: "main_convert-by-profile_no-wdir-dot_02",
        IsMulti:  true,
        InputArr: []string{
          "shapes",         // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_TXT_SHAPES_ONLY,
      },
      {
        TestName: "main_convert-by-profile_confusing-root-path_00",
        IsMulti:  true,
        InputArr: []string{
          "./fruits/..",    // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_ALL,
      },
      {
        TestName: "main_convert-by-profile_confusing-root-path_01",
        IsMulti:  true,
        InputArr: []string{
          "./ABC/..",       // path to working directory.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_ALL,
      },
      {
        TestName: "main_convert-by-profile_wdir-not-a-dir_00",
        IsMulti:  true,
        InputArr: []string{
          "fruits/bananas.txt", // path to working directory.
          "config",             // config path.
          "general.config",     // config name.
          "prettify-txt",       // profile name.
          BASIC_CONF,           // config content.
        },
        ExpectedValue: FIND_TXT_BANANAS,
      },
      {
        TestName: "main_convert-by-profile_wdir-not-a-dir_01",
        IsMulti:  true,
        InputArr: []string{
          "fruits/bananas",     // path to working directory.
          "config",             // config path.
          "general.config",     // config name.
          "prettify-txt",       // profile name.
          BASIC_CONF,           // config content.
        },
        ExpectedValue: FIND_TXT_BANANAS,
      },
    },
  },

  /*
  * NOTE: this actually should be a E2E test.
  * Pipe test: ls   | grep -E 'txt$' | renamer -profile files_txt
  * Pipe test: find | grep -E 'txt$' | renamer -profile files_txt
  */
  {
    TestingFunction:
    func(t *testing.T, in gt.TestList) string {
      // set test variables.
      firstCommand  := in.InputArr[0]
      configPath    := in.InputArr[1]
      configName    := in.InputArr[2]
      profileName   := in.InputArr[3]
      configContent := in.InputArr[4]
      // run test setup.
      path := tu.MakeRealTestFs()
      tu.CreateFile(path, configPath, configName, configContent)
      // simulate pipe.
      cmds := CommandList{
        {
          firstCommand,
          []string{},
        },
        {
          "grep",
          []string{"-E", ".txt$"},
        },
      }
      inputListing := simulatePipe(cmds, path, t)
      // create test file system.
      mdir          := tu.ManageDir()
      mdir.FillFile(configPath, configName, configContent)
      conf          := mdir.SubPath(configPath + "/" + configName)
      // main test.
      _ = ConvertByProfile( // profile command: "renamer -profile 'profileName'"
        mdir.FsSub,         // file system.
        ".",                // path to working directory.
        conf,               // path to config.
        profileName,        // profile.
        inputListing,       // input.
      )
      ExecuteByApplying()
      // remove test file system.
      outputListing := mdir.ListTreeOsfs()
      mdir.CleanUp()
      // return.
      return outputListing
    },
    Tests:
    []gt.TestList{
      {
        TestName: "full-test_pipe-test_00",
        IsMulti:  true,
        InputArr: []string{
          "ls",             // first command.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: LS_TXT,
      },
      {
        TestName: "full-test_pipe-test_00",
        IsMulti:  true,
        InputArr: []string{
          "find",           // first command.
          "config",         // config path.
          "general.config", // config name.
          "prettify-txt",   // profile name.
          BASIC_CONF,       // config content.
        },
        ExpectedValue: FIND_TXT,
      },
    },
  },

  /* Fin test suite. */
}

