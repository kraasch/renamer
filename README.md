
# !!! UNFINISHED WORK !!!

This program doesn't work and can delete files.

## renamer

## overview

All source files and their packages:

```text
.
├── cmd
│   └── renamer.go // CLI program
└── pkg
    ├── edit       // calls editor.
    ├── fsmanage   // deals with file system.
    ├── dir        // lists directories.
    ├── testutil   // creates mock file system.
    ├── rnmanage   // TODO: ...
    ├── autorn     // TODO: ...
    └── rename     // renames things.
```

Dependencies of the CLI program `renamer`:

```text
renamer
├── edit
└── rnmanage
```

Dependencies of the core package `rnmanage`:

```text
rnmanage
├── fsmanage
│   ├── dir
│   └── testutil
└── autorn
    └── rename
```

## tasks

  - [X] use this for filesystem mocks: https://github.com/spf13/afero

Idea: validate file names. (version 1)

```go
import (
	"unicode"
)
func isValidFileName(name string) bool {
	// Define a range of allowed characters for a valid filename
	for _, r := range name {
		if !unicode.IsPrint(r) || r == '/' || r == '\\' || r == ':' || r == '*' || r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return false
		}
	}
	return true
}
```

Idea: validate file names. (version 2)

```go
package main

import (
	"unicode"
)

func isValidFileName(name string) bool {
	// Check that each character is within the allowed Unicode range for filenames
	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.' && r != '-' && r != '_' {
			return false
		}
	}
	return true
}
```

## tests

Create tests:

 - [ ] The program should not delete files by accident by moving or renaming them into already existing files.
 - [ ] The program should be able to rename two files into each other.


Note: The directory could be this, for example...

```text
a
b
```

but the edited lines could be

```text
b
a
```


In this case both files collide and cannot be renamed unless the other file has been renamed.
The program should be able to deal with some conflicts, either by having an intermediate name or detecting such collisions and renaming in a smart way.
