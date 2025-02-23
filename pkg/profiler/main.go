
package profiler

import (
  // other packages.
  toml "github.com/BurntSushi/toml"

  // other packages.
  "bytes"

  // local packages.
  rn "github.com/kraasch/renamer/pkg/rename"
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

func (c *Config) AddProfileToConfig(p *Profile, name string) Config {
  c.Profiles[name] = p
  return *c
}

func (p *Profile) Apply(input string) string {
  r := p.ProfileRule
  return rn.ApplyRenamingRules(
    input,
    r.WordSeparators,
    r.DeleteChars,
    r.Conversions,
    r.SmallGapMark,
    r.BigGapMark,
    r.ModesString,
  )
}

func (c *Config) ToToml() string {
  var buf bytes.Buffer
  if err := toml.NewEncoder(&buf).Encode(c); err != nil {
    panic(err)
  }
  return buf.String()
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

