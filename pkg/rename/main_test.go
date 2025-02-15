
package rename

import (
  "testing"
  gt "github.com/kraasch/gotest/gotest"
)

func TestAll(t *testing.T) {
  gt.DoTest(t, suites)
}

var suites = []gt.TestSuite{
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
    TestingFunction:
    func(in gt.TestList) (out string) {
      wordSeparators := in.InputArr[0]
      deleteChars    := in.InputArr[1]
      smallGapMark   := in.InputArr[2]
      bigGapMark     := in.InputArr[3]
      conversions    := in.InputArr[4]
      modesString    := in.InputArr[5]
      targetName     := in.InputArr[6]
      out = ApplyRenamingRules(targetName, wordSeparators, deleteChars, conversions, smallGapMark, bigGapMark, modesString)
      return
    },
    Tests:
    []gt.TestList{
      {
        TestName: "rename-file_common-file-name_00",
        IsMulti:  false,
        InputArr: []string{
          " ()", // word separators.
          "",    // delete characters.
          "-",   // small gap replacement.
          "_",   // big gap replacement.
          "Aa",  // list of conversions.
          "",    // string of modes.
          "The Walking Dead S05E01 No Sanctuary (1080p x265 Joy).mkv",
        },
        ExpectedValue:
        "the-walking-dead-s05e01-no-sanctuary_1080p-x265-joy.mkv",
      },
      {
        TestName: "rename-file_common-file-name_01",
        IsMulti:  false,
        InputArr: []string{
          " ():", // word separators.
          "'",    // delete characters.
          "-",    // small gap replacement.
          "_",    // big gap replacement.
          "Aa",   // list of conversions.
          "",     // string of modes.
          "Head First Software Architecture: A Learner's Guide to Architectural Thinking (English Edition--).pdf",
        },
        ExpectedValue:
        "head-first-software-architecture_a-learners-guide-to-architectural-thinking_english-edition.pdf",
      },
      {
        TestName: "rename-file_common-file-name_02",
        IsMulti:  false,
        InputArr: []string{
          " ():", // word separators.
          ",.",   // delete characters.
          "-",    // small gap replacement.
          "_",    // big gap replacement.
          "Aa",   // list of conversions.
          "",     // string of modes.
          "The Internal-Combustion Engine in Theory and Practice. Vol. I Thermodynamics, Fluid Flow, Performance ( PDFDrive ).pdf",
        },
        ExpectedValue:
        "the-internal-combustion-engine-in-theory-and-practice-vol-i-thermodynamics-fluid-flow-performance_pdfdrive.pdf",
      },
      {
        TestName: "rename-file_mode_replace-last-dot_00",
        IsMulti:  false,
        InputArr: []string{
          " ():", // word separators.
          ",.",   // delete characters.
          "-",    // small gap replacement.
          "_",    // big gap replacement.
          "Aa",   // list of conversions.
          "D",    // string of modes.
          "The Internal-Combustion Engine in Theory and Practice. Vol. I Thermodynamics, Fluid Flow, Performance ( PDFDrive ).pdf",
        },
        ExpectedValue:
        "the-internal-combustion-engine-in-theory-and-practice-vol-i-thermodynamics-fluid-flow-performance_pdfdrive_pdf",
      },
    },
  },
  /* Fin test suite. */
}

