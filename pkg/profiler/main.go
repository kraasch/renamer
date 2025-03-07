
package profiler

import (
  // other packages.
  toml "github.com/BurntSushi/toml"

  // other packages.
  "bytes"
  "strings"
  "fmt"
  "slices"

  // local packages.
  rn "github.com/kraasch/renamer/pkg/rename"
)

const (
  PIPE  = "|"
  BEG   = "【"
  END   = "】"
)

type Config struct {
  Title    string              `toml:"title"`
  Profiles map[string]*Profile `toml:"profiles"`
}

type Profile struct {
  Name        string           `toml:"name"`
  ProfileRule Rule             `toml:"profile_rule"`
}

type Rule struct {
  WordSeparators string        `toml:"word_separators"`
  DeleteChars    string        `toml:"delete_chars"`
  SmallGapMark   string        `toml:"small_gap_mark"`
  BigGapMark     string        `toml:"big_gap_mark"`
  Conversions    string        `toml:"conversions"`
  ModesString    string        `toml:"modes_string"`
}

func (r Rule) String() string {
  return fmt.Sprintf("%s%s|%s|%s|%s|%s|%s%s", BEG, r.WordSeparators, r.DeleteChars, r.SmallGapMark, r.BigGapMark, r.Conversions, r.ModesString, END)
}

func (c Config) String() string {
  res := []string{}
  profiles := c.Profiles
  maxNameLen := 0
  maxIdLen   := 0
  for pid := range profiles { // get length of longest profile id.
    idl   := len(fmt.Sprintf("%+v\n", pid))
    namel := len(profiles[pid].Name)
    if idl > maxIdLen {
      maxIdLen = idl
    }
    if namel > maxNameLen {
      maxNameLen = namel
    }
  }
  longestLineLen := 0
  for pid := range profiles {
    pname := profiles[pid].Name
    prule := profiles[pid].ProfileRule
    formatted := fmt.Sprintf("%+-*v(%-*s) %s\n", maxIdLen, pid, maxNameLen, pname, prule)
    res = append(res, formatted)
    // see which line is the longest.
    if len(formatted) > longestLineLen {
      longestLineLen = len(formatted)
    }
  }
  // sort and format.
  slices.Sort(res)
  out := c.Title + ":\n"
  out += strings.Repeat("-", longestLineLen) + "\n"
  for _, r := range res {
    out += r
  }
  return out
}

func (c *Config) AddProfileToConfig(p *Profile, name string) Config {
  c.Profiles[name] = p
  return *c
}

func (p *Profile) Apply(input string, metaInfo rn.MetaInfo) string {
  r := p.ProfileRule
  return rn.ApplyRenamingRules(
    input,
    r.WordSeparators,
    r.DeleteChars,
    r.Conversions,
    r.SmallGapMark,
    r.BigGapMark,
    r.ModesString,
    metaInfo,
  )
}

func (c *Config) ToToml() string {
  var buf bytes.Buffer
  if err := toml.NewEncoder(&buf).Encode(c); err != nil {
    panic(err)
  }
  return buf.String()
}

func ProfileFromRuleString(ruleString string) Profile {
  // decode rule string.
  ss := strings.Split(ruleString, PIPE) // TODO: throw error if split is too long or short.
  rule := Rule {
    WordSeparators: ss[0],
    DeleteChars:    ss[1],
    SmallGapMark:   ss[2],
    BigGapMark:     ss[3],
    Conversions:    ss[4],
    ModesString:    ss[5],
  }
  profile := Profile{
    Name:        "converted from rule #000", // TODO: increase number.
    ProfileRule: rule,
  }
  return profile
}

func ReadRawProfileConfig(tomlBlob string) Config {
  // decode toml.
  var c Config
	_, err := toml.Decode(tomlBlob, &c)
	if err != nil {
		panic("Failed to decode TOML.")
	}
  return c
}

