
package rname

import (

  // this is a test.
  "testing"

  // printing and formatting.
  "fmt"

  // other imports.
  "github.com/kraasch/godiff/godiff"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  isMulti           bool
  inputArr          []string
  expectedValue     string
}

type TestSuite struct {
  testingFunction   func(in TestList) string
  tests             []TestList
}

var suites = []TestSuite{
  // /*
  // * Test error states for the ApplyRenamingRules().
  // */
  // {
  //   testingFunction:
  //   func(in TestList) (string) {
  //     wordSeparators := in.inputArr[0]
  //     deleteChars    := in.inputArr[1]
  //     smallGapMark   := in.inputArr[2]
  //     bigGapMark     := in.inputArr[3]
  //     conversions    := in.inputArr[4]
  //     modesString    := in.inputArr[5]
  //     targetName     := in.inputArr[6]
  //     out, err := ApplyRenamingRules(targetName, wordSeparators, deleteChars, conversions, smallGapMark, bigGapMark, modesString)
  //     if (err != nil) {
  //        out = fmt.Sprint(err) // successfully received and error, thus compare error message string.
  //     }
  //     t.Errorf("Dit not receive any error.\n")// fail here.
  //     return out
  //   },
  //   tests:
  //   []TestList{
  //     {
  //       testName: "rename-file_error_has-pipe_00",
  //       isMulti:  false,
  //       inputArr: []string{
  //         " ()",  // word separators.
  //         "",     // delete characters.
  //         "-",    // small gap replacement.
  //         "_",    // big gap replacement.
  //         "Aa",   // list of conversions.
  //         "",     // string of modes.
  //         "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
  //       },
  //       expectedValue: "Input parameters cannot contain pipes", // Error message.
  //     },
  //   },
  // },
  /*
  * Test for the ApplyRenamingRules().
  */
  {
    testingFunction:
    func(in TestList) (out string) {
      wordSeparators := in.inputArr[0]
      deleteChars    := in.inputArr[1]
      smallGapMark   := in.inputArr[2]
      bigGapMark     := in.inputArr[3]
      conversions    := in.inputArr[4]
      modesString    := in.inputArr[5]
      targetName     := in.inputArr[6]
      out = ApplyRenamingRules(targetName, wordSeparators, deleteChars, conversions, smallGapMark, bigGapMark, modesString)
      return
    },
    tests:
    []TestList{
      {
        testName: "rename-file_common-file-name_00",
        isMulti:  false,
        inputArr: []string{
          " ()", // word separators.
          "",    // delete characters.
          "-",   // small gap replacement.
          "_",   // big gap replacement.
          "Aa",  // list of conversions.
          "",    // string of modes.
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        expectedValue:
        "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        testName: "rename-file_common-file-name_01",
        isMulti:  false,
        inputArr: []string{
          " ():", // word separators.
          "'",    // delete characters.
          "-",    // small gap replacement.
          "_",    // big gap replacement.
          "Aa",   // list of conversions.
          "",     // string of modes.
          "Head First Software Architecture: A Learner's Guide to Architectural Thinking (English Edition--).pdf",
        },
        expectedValue:
        "head-first-software-architecture_a-learners-guide-to-architectural-thinking_english-edition.pdf",
      },
      {
        testName: "rename-file_common-file-name_02",
        isMulti:  false,
        inputArr: []string{
          " ():", // word separators.
          ",.",   // delete characters.
          "-",    // small gap replacement.
          "_",    // big gap replacement.
          "Aa",   // list of conversions.
          "",     // string of modes.
          "The Internal-Combustion Engine in Theory and Practice. Vol. I Thermodynamics, Fluid Flow, Performance ( PDFDrive ).pdf",
        },
        expectedValue:
        "the-internal-combustion-engine-in-theory-and-practice-vol-i-thermodynamics-fluid-flow-performance_pdfdrive.pdf",
      },
      {
        testName: "rename-file_mode_replace-last-dot_00",
        isMulti:  false,
        inputArr: []string{
          " ():", // word separators.
          ",.",   // delete characters.
          "-",    // small gap replacement.
          "_",    // big gap replacement.
          "Aa",   // list of conversions.
          "D",    // string of modes.
          "The Internal-Combustion Engine in Theory and Practice. Vol. I Thermodynamics, Fluid Flow, Performance ( PDFDrive ).pdf",
        },
        expectedValue:
        "the-internal-combustion-engine-in-theory-and-practice-vol-i-thermodynamics-fluid-flow-performance_pdfdrive_pdf",
      },
    },
  },
  /* Fin test suite. */
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      t.Run(name, func(t *testing.T) {
        exp := test.expectedValue
        got := suite.testingFunction(test)
        if exp != got {
          if test.isMulti {
            t.Errorf("In '%s':\n", name)
            diff := godiff.CDiff(exp, got)
            t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
            t.Errorf("exp/got:\n%s\n", diff)
          } else {
            t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
          }
        }
      })
    }
  }
}

