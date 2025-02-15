
package rename

import (
  "regexp"
  "strings"
)

const (
  sectionDot = "."
  sectionSep = "|"
)

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

  // apply conversions.
  {
    arr := strings.Split(conversions, ",")
    for _, conversion := range arr {
      from := conversion[0] // first rune.
      into := conversion[1] // second rune.
      if from == 'A' && into == 'a' {
        s = strings.ToLower(s)
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

