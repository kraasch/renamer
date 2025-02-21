
package rename

import (
  "regexp"
  "strings"
  "unicode"
)

const (
  sectionDot = "."
  sectionSep = "|"
)

func DeleteNonAscii(s string) string {
  var result strings.Builder
  for _, r := range s {
    if r < unicode.MaxASCII {
      result.WriteRune(r)
    }
  }
  return result.String()
}

func replaceLast(target, from, into string) (result string) {
  i := strings.LastIndex(target, from)
  if i == -1 {
    return target
  }
  return target[:i] + into + target[i+len(from):]
}

func ApplyRenamingRules(targetName, wordSeparators, deleteChars, conversions, smallGapMark, bigGapMark, modes string) (s string) {

  // start with string to rename.
  s = targetName

  // read modes.
  deleteLastDot := false
  for _, mode := range modes {
    if mode == 'D' {
      deleteLastDot = true
    }
  }

  // modes.
  // mode: preserve last dot.
  if !deleteLastDot {
    s = replaceLast(s, sectionDot, sectionSep) // temporarily replace last dot with pipe.
  }

  // delete characters.
  for _, char := range deleteChars {
    s = strings.ReplaceAll(s, string(char), "")
  }

  // find word separators.
  for _, char := range wordSeparators {
    s = strings.ReplaceAll(s, string(char), smallGapMark)
  }

  // apply actions.
  {
    actions := []struct {
      Mnemonic string
      Function   func(string) string
    }{

      // conversions.
      {"cAa", strings.ToLower}, // convert capital to lower.
      //{"caA", strings.ToUpper}, // convert lower to capital.

      // deletions.
      {"dna", DeleteNonAscii},  // delete non-ascii characters.

    }
    arr := strings.Split(conversions, ",")
    for _, actionStr := range arr {
      for _, action := range actions {
        if actionStr == action.Mnemonic {
          s = action.Function(s)
        }
      }
    }
  }

  // find groups of small marks and make them into big marks.
  {
    re, _ := regexp.Compile(smallGapMark + "[" + smallGapMark + "]+")
    s = re.ReplaceAllString(s, bigGapMark)
  }

  // treat dot (.) as section separator and swallow other gap markers that appear around it.
  if !deleteLastDot {
    s = replaceLast(s, sectionSep, sectionDot) // put dot back in.
  }
  {
    re, _ := regexp.Compile("[" + bigGapMark + smallGapMark + "]*\\.[" + bigGapMark + smallGapMark + "]*")
    s = re.ReplaceAllString(s, ".")
  }

  return
}

