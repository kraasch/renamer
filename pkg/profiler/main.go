
package profiler

import (
  // other packages.
  toml "github.com/BurntSushi/toml"

  // other packages.
  // "fmt"
)

type Config struct {
  Title    string
  Profiles map[string]Profile
}

type Profile struct {
  Name string
  Rule string
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


