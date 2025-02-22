
package rename

import (
  "regexp"
  "strings"
  "unicode"
)

var (
  actions = []struct {
    Mnemonic    string
    Description string
    Function    func(string) string
  }{

    /*
    * conversions.
    */
    {
      "cAa", "convert capital to lower (but not file ending)",
      func(s string) string {
        before := s
        ending := ""
        lastIndex := strings.LastIndex(s, ".")
        if lastIndex != -1 {
          before = s[:lastIndex]
          ending = s[lastIndex:]
        }
        return strings.ToLower(before) + ending
      },
    },

    {
      "caA", "convert lower to capital (but not file ending)",
      func(s string) string {
        before := s
        ending := ""
        lastIndex := strings.LastIndex(s, ".")
        if lastIndex != -1 {
          before = s[:lastIndex]
          ending = s[lastIndex:]
        }
        return strings.ToUpper(before) + ending
      },
    },

    {
      "CAa", "convert capital to lower (including file ending)",
      func(s string) string {
        return strings.ToLower(s)
      },
    },

    {
      "CaA", "convert lower to capital (including file ending)",
      func(s string) string {
        return strings.ToUpper(s)
      },
    },

    /*
    * deletions.
    */
    {
      "dna", "delete non-ascii characters",
      func(s string) string {
        var result strings.Builder
        for _, r := range s {
          if r < unicode.MaxASCII {
            result.WriteRune(r)
          }
        }
        return result.String()
      },
    },

    {
      "dnr", "delete non-readable characters (not letters, digits or in {.-_})",
      func(s string) string {
        var result strings.Builder
        for _, r := range s {
          if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '-' || r == '_' {
            result.WriteRune(r)
          }
        }
        return result.String()
      },
    },

    /*
    * insertions.
    */
    {
      "id^", "insert current date in beginning",
      func(s string) string {
        var result strings.Builder
        date := "2020-12-20" // TODO: get current date.
        date += "_"
        for _, r := range date {
          result.WriteRune(r)
        }
        for _, r := range s {
          result.WriteRune(r)
        }
        return result.String()
      },
    },

    {
      "id$", "insert current date in end",
      func(s string) string {
        var result strings.Builder
        date := "_"
        date += "2020-12-20" // TODO: get current date.
        for _, r := range s {
          result.WriteRune(r)
        }
        for _, r := range date {
          result.WriteRune(r)
        }
        return result.String()
      },
    },

    {
      "id.", "insert current date before file ending",
      func(s string) string {
        before := s
        ending := ""
        lastIndex := strings.LastIndex(s, ".")
        if lastIndex != -1 {
          before = s[:lastIndex]
          ending = s[lastIndex:]
        }
        var result strings.Builder
        date := "_"
        date += "2020-12-20" // TODO: get current date.
        for _, r := range before {
          result.WriteRune(r)
        }
        for _, r := range date {
          result.WriteRune(r)
        }
        for _, r := range ending {
          result.WriteRune(r)
        }
        return result.String()
      },
    },

    // Fin of actions.
  }
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

func ValidateRenamingRules(targetName, wordSeparators, deleteChars, conversions, smallGapMark, bigGapMark, modes string) (bool) {
  resultOfApplyingRules := ApplyRenamingRules(targetName, wordSeparators, deleteChars, conversions, smallGapMark, bigGapMark, modes)
  return targetName == resultOfApplyingRules
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
  if smallGapMark != "" && bigGapMark != "" {
    if !deleteLastDot {
      s = replaceLast(s, sectionDot, sectionSep) // temporarily replace last dot with pipe.
    }
  }

  // delete characters.
  for _, char := range deleteChars {
    s = strings.ReplaceAll(s, string(char), "")
  }

  // find word separators.
  if smallGapMark != "" {
    for _, char := range wordSeparators {
      s = strings.ReplaceAll(s, string(char), smallGapMark)
    }
  }

  // apply actions.
  {
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
  if smallGapMark != "" && bigGapMark != "" {
    re, _ := regexp.Compile(smallGapMark + "[" + smallGapMark + "]+")
    s = re.ReplaceAllString(s, bigGapMark)
  }

  // treat dot (.) as section separator and swallow other gap markers that appear around it.
  if smallGapMark != "" && bigGapMark != "" {
    if !deleteLastDot {
      s = replaceLast(s, sectionSep, sectionDot) // put dot back in.
    }
    {
      re, _ := regexp.Compile("[" + bigGapMark + smallGapMark + "]*\\.[" + bigGapMark + smallGapMark + "]*")
      s = re.ReplaceAllString(s, ".")
    }
  }

  return
}

