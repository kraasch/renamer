//go:build testing

package testutil

import (
  "fmt"
  "github.com/spf13/afero"
  // "io/fs"
)

func MakeTestFs() afero.Fs {
  var fileSystem = afero.NewMemMapFs()
  dirs := []string{
    "fruits/",
    "shapes/",
  }
  files := []string{
    "notes.txt",
    "fruits/apples.txt",
    "fruits/bananas.txt",
    "shapes/triangle.txt",
    "fruits/coconuts.txt",
    "shapes/square.txt",
    "shapes/circle.txt",
  }
  for _, dir := range dirs {
    if err := fileSystem.MkdirAll(dir, 0755); err != nil {
      fmt.Println("Setting up test failed.") // TODO: implement test failure.
      return nil
    }
  }
  for _, file := range files {
    if err := afero.WriteFile(fileSystem, file, []byte("Not empty."), 0644); err != nil {
      fmt.Println("Setting up test failed.") // TODO: implement test failure.
      return nil
    }
  }
  // return afero.NewIOFS(fileSystem)
  return fileSystem
}

