//g o:build testing // TODO: comment in again.

package testutil

import (
  // create fake file system.
  "fmt"
  "github.com/spf13/afero"
  // create real file system.
  "os"
  "path/filepath"
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
  // return afero.NewIOFS(fileSystem) // NOTE: parses afero to os.FS.
  return fileSystem
}

func MakeRealTestFs() string {
  testdir := "testfs"
  newpath := filepath.Join(".", testdir)
  // delete any old remains of test directory.
  os.RemoveAll(newpath) // NOTE: recursively deletes path.
  // create test directory.
  if _, err := os.Stat(newpath); err == nil  {
    // directory already existed, therefore stop testing.
    panic("TEST SETUP: Directory 'testfs' already exists.")
  }
  if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
    panic("TEST SETUP: Directory 'testfs' could not be made.")
  }
  // create test dir tree.
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
    dirPath := filepath.Join(".", testdir, dir)
    if err := os.MkdirAll(dirPath, 0755); err != nil {
      panic("TEST SETUP: Creation of directory tree failed.")
    }
  }
  for _, file := range files {
    filePath := filepath.Join(".", testdir, file)
    if err := os.WriteFile(filePath, []byte("Not empty."), 0644); err != nil {
      panic("TEST SETUP: Creation of file tree failed.")
    }
  }
  return string(newpath)
}

func CleanUpRealTestFs(subdir string) {
  path := filepath.Join(".", subdir)
  // delete any old remains of test directory.
  os.RemoveAll(path) // NOTE: recursively deletes path.
  // TODO: report failure.
}

