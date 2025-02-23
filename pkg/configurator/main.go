
package configurator

import (
  "os"
  "fmt"
)

func ReadConfig(configPath string) string {
  dat, err := os.ReadFile("." + "/" + configPath)
  if err != nil {
    panic(err)
  }
  return string(dat)
}

func Toast(in string) string {
  return fmt.Sprintf("This is %s!\n", in)
}

