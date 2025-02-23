
package edit

import (
  "errors"
  // "fmt"
  // "bufio"
  // "os"
  // "os/exec"
  // "strings"
  // "log"
)

type Editor struct {
  editHistory []func(string) (string, error)
  i           int
}

func NewEditor(hist []func(string) (string, error)) Editor {
  return Editor{
    hist,
    0,
  }
}

func (e *Editor) Edit(in string) (string, error) {
  i := e.i
  if i < len(e.editHistory) {
    return e.editHistory[i](in)
  }
  return "", errors.New("Edit history exhausted.")
}

/* TODO: make this into a sub module, ie package. */

// func manualRename() {
//   // get the editor from the environment.
//   editor := os.Getenv("EDITOR")
//   if editor == "" {
//     log.Fatal("EDITOR environment variable is not set.")
//   }
//   fmt.Printf("Using editor '%s'.\n", editor)
// 
//   // get the current directory.
//   dir, err := os.Getwd()
//   if err != nil {
//     log.Fatal(err)
//   }
//   fmt.Printf("Renaming directory '%s'.\n", dir)
// 
//   // read the files in the current directory.
//   files, err := os.ReadDir(dir)
//   if err != nil {
//     log.Fatal(err)
//   }
// 
//   // filter the list of files to only include regular files (ignore directories).
//   var fileNames []string
//   for _, file := range files {
//     if !file.IsDir() {
//       fileNames = append(fileNames, file.Name())
//     }
//   }
//   fmt.Printf("Files before: '%s'.\n", fileNames)
// 
//   // open an editor to edit the list of files.
//   fileListFile, err := os.CreateTemp("", "file_list_*.txt")
//   if err != nil {
//     log.Fatal(err)
//   }
//   defer os.Remove(fileListFile.Name()) // cleanup the temp file.
// 
//   // write the file names to the temp file, one per line.
//   writer := bufio.NewWriter(fileListFile)
//   for _, fileName := range fileNames {
//     writer.WriteString(fileName + "\n") // TODO: check error.
//   }
//   writer.Flush()
// 
//   // open the editor with the temp file.
//   cmd := exec.Command(editor, fileListFile.Name())
//   cmd.Stdout = os.Stdout
//   cmd.Stderr = os.Stderr
//   if err := cmd.Run(); err != nil {
//     log.Fatalf("Error while launching the editor: %v", err)
//   }
// 
//   // read the edited list of file names.
//   fileListFile.Seek(0, 0) // TODO: check error.
//   scanner := bufio.NewScanner(fileListFile)
//   editedFileNames := []string{}
//   intermFileNames := []string{}
//   for scanner.Scan() {
//     editedFileNames = append(editedFileNames, strings.TrimSpace(scanner.Text()))
//   }
//   if err := scanner.Err(); err != nil {
//     log.Fatal(err)
//   }
//   fmt.Printf("Files after: '%s'.\n", editedFileNames)
// 
//   // handle renaming files carefully to avoid collisions.
//   renamedFiles := make(map[string]string)
//   for i, oldName := range fileNames {
//     newName := editedFileNames[i]
//     if newName == oldName {
//       fmt.Printf("Nothing to do: %s == %s.\n", oldName, newName)
//       continue
//     }
// 
//     // check for filename collisions.
//     if _, exists := renamedFiles[newName]; exists {
//       // there is a collision, rename with an intermediate name.
//       someHash := "b8b81ed4" // TODO: make this random
//       intermediateName := fmt.Sprintf("%s_%d_%s", newName, i, someHash)
//       fmt.Printf("Renaming %s to %s (intermediate)\n", oldName, intermediateName)
//       if err := os.Rename(oldName, intermediateName); err != nil {
//         log.Fatal(err)
//       } else {
//         intermFileNames = append(intermFileNames, intermediateName)
//         fmt.Printf(" 1 - Renamed: %s to %s.\n", oldName, intermediateName)
//       }
//       renamedFiles[intermediateName] = oldName
//       continue
//     }
// 
//     // rename the file if there's no conflict.
//     if err := os.Rename(oldName, newName); err != nil {
//       log.Fatal(err)
//     } else {
//       fmt.Printf(" 2 - Renamed: %s to %s.\n", oldName, newName)
//     }
//     renamedFiles[newName] = oldName
//   }
// 
//   // ensure that files that were renamed to intermediate names are renamed to their final names.
//   for _, intermediateName := range intermFileNames {
//     finalName := renamedFiles[intermediateName]
//     if intermediateName != finalName {
//       if err := os.Rename(intermediateName, finalName); err != nil {
//         log.Fatal(err)
//       } else {
//         fmt.Printf(" 3 - Renamed: %s to %s.\n", intermediateName, finalName)
//       }
//     }
//   }
// 
//   fmt.Println("Files renamed successfully.")
// }
