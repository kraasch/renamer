//g o:build testing // TODO: comment in again.

package testutil

import (
  // create fake file system.
  "fmt"
  "github.com/spf13/afero"

  // create real file system.
  "os"
  "path/filepath"

  // list files in directory.
  "strings"
  iofs "io/fs"
)

const (
  DIRSPERM = 0755
  FILEPERM = 0644
)

func ListFs(fs afero.Fs, path string) string {
  var builder strings.Builder
  _ = afero.GetTempDir(fs, path)
  err := afero.Walk(fs, path, func(filePath string, info iofs.FileInfo, err error) error {
    if err != nil {
      return err
    }
    builder.WriteString(filePath + "\n")
    return nil
  })
  if err != nil {
    panic(err)
  }
  return builder.String()
}

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
    if err := fileSystem.MkdirAll(dir, DIRSPERM); err != nil {
      fmt.Println("Setting up test failed.") // TODO: implement test failure.
      return nil
    }
  }
  for _, file := range files {
    if err := afero.WriteFile(fileSystem, file, []byte("Not empty."), FILEPERM); err != nil {
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
    if err := os.MkdirAll(dirPath, DIRSPERM); err != nil {
      panic("TEST SETUP: Creation of directory tree failed.")
    }
  }
  for _, file := range files {
    path := filepath.Join(".", testdir, file)
    if err := os.WriteFile(path, []byte("Not empty."), FILEPERM); err != nil {
      panic("TEST SETUP: Creation of file tree failed.")
    }
  }
  return string(newpath)
}

func CreateFile(testDir, fileDir, fileName, fileContent string) {
  // TODO: Generallize, so far this function does not create file paths,
  // can only create one subfolder (depth level = 1).
  path := filepath.Join(".", testDir, fileDir)
  fmt.Println("Try to create file at", path)
  if err := os.MkdirAll(path, DIRSPERM); err != nil {
    fmt.Println("NOPE 1")
    {}
  }
  path = filepath.Join(".", testDir, fileDir, fileName)
  if err := os.WriteFile(path, []byte(fileContent), FILEPERM); err != nil {
    fmt.Println("NOPE 2")
    {}
  }
  // TODO: report failure.
}

func CleanUpRealTestFs(testDir string) {
  path := filepath.Join(".", testDir)
  // delete any old remains of test directory.
  os.RemoveAll(path) // NOTE: recursively deletes path.
  // TODO: report failure.
}

