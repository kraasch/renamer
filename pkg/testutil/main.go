//g o:build testing // TODO: comment in again.

package testutil

import (
  // create fake file system.
  "fmt"
  afero "github.com/spf13/afero"
  iofs "io/fs"

  // create real file system.
  "os"
  "path/filepath"

  // list files in directory.
  "strings"
  "sort"
)

const (
  DIRSPERM = 0755
  FILEPERM = 0644
)

type ManagedDir struct {
  path       string
  FsOriginal afero.Fs
  FsSub      afero.Fs
}

func ManageDir() ManagedDir {
  var md ManagedDir
  md.path = MakeRealTestFs()
  // init fs.
  md.FsOriginal  = afero.NewOsFs()
  currentDir, _ := os.Getwd()
  targetDir     := filepath.Join(currentDir, "testfs")
  md.FsSub       = afero.NewBasePathFs(md.FsOriginal, targetDir)
  // return.
  return md
}

func (md *ManagedDir) FillFile(path, name, content string) {
  CreateFile(md.path, path, name, content)
}

func (md *ManagedDir) CleanUp() {
  CleanUpRealTestFs(md.path)
}

func (md *ManagedDir) ListTree() string {
  listing       := DirListTree(afero.NewIOFS(md.FsSub))
  return listing
}

/*
* NOTE: essentially does the same as ListTree().
*/
func (md *ManagedDir) ListTreeOsfs() string {
  fs2     := afero.NewIOFS(md.FsOriginal)
  fs3, _  := fs2.Sub("testfs")
  listing := DirListTree(fs3)
  return listing
}

func (md *ManagedDir) SubPath(path string) string {
      return "testfs/" + path
}

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

func MakeEmptyTestFs() afero.Fs {
  return MakeFs([]string{}, []string{})
}

func MakeTestFs() afero.Fs {
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
  return MakeFs(dirs, files)
}

func MakeFs(dirs, files []string) afero.Fs {
  var fileSystem = afero.NewMemMapFs()
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
  return MakeRealFs(dirs, files)
}

func MakeEmptyRealTestFs() string {
  return MakeRealFs([]string{},[]string{})
}

func MakeRealFs(dirs, files []string) string {
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

func DirList(fileSystem iofs.FS) string {
  dirEntries, err := iofs.ReadDir(fileSystem, ".")
  if err != nil {
    return "" // TODO: handle errors appropriately.
  }
  var entries []string
  for _, entry := range dirEntries {
    name := entry.Name()
    if entry.IsDir() {
      name += "/" // add trailing slash for directories.
    }
    entries = append(entries, name)
  }
  out := strings.Join(entries, "\n")
  return out
}

func DirListTree(fileSystem iofs.FS) string {
  var fileList []string
  err := iofs.WalkDir(fileSystem, ".", func(path string, d iofs.DirEntry, err error) error {
    if err != nil {
      return err // TODO: handle errors appropriately.
    }
    if path == "." {
      return nil // TODO: skip the root directory itself.
    }
    if d.IsDir() {
      fileList = append(fileList, path + "/") // add trailing slash for directories.
    } else {
      fileList = append(fileList, path)
    }
    return nil
  })
  if err != nil {
      return "" // TODO: handle errors appropriately.
  }
  sort.Strings(fileList) // sort alphabetically.
  out := strings.Join(fileList, "\n")
  return out
}

