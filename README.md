
# !!! UNFINISHED WORK !!!

This program doesn't work and can delete files.

## renamer

## overview

General structure of source files and their packages:

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
    ├── profiler   // TODO: ...
    └── rename     // renames strings.
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
    ├── profiler
    └── rename
```

## tasks

Main functionality:

  - INPUT.
    - [ ] get file list from pipe, ie allow these inputs:
      - [ ] `ls | grep -E 'mp3$' | renamer -edit`
      - [ ] `find | grep -E '.ogg$' | renamer -profile music_ogg`
    - [ ] if nothing is piped in, have two search modes: `recusvie` and `list`
      - [ ] `renamer -list=all   -profile music_ogg` list all in pwd (default).
      - [ ] `renamer -list=files -profile music_ogg` list all files in pwd.
      - [ ] `renamer -list=dirs  -profile music_ogg` list all directories in pwd.
      - [ ] `renamer -recursive=all   -profile music_ogg` as above, but recursive (default)
      - [ ] `renamer -recursive=files -profile music_ogg` as above, but recursive.
      - [ ] `renamer -recursive=dirs  -profile music_ogg` as above, but recursive.
  - RENAMING.
    - [ ] different ways of renaming.
      - [ ] -edit opens editor
      - [ ] -auto just applies the profile
      - [ ] -interactive lets user choose a profile for each file or manually edit it
    - [ ] implement renaming profiles.
      - [ ] run `renamer -profile media` to rename media files with specified
            conversion (named `media`, for renaming music, video, etc).
      - [ ] default to `renamer -profile default` if no profile is specified.
    - [ ] implement pure editor version (use native editor, eg vi, emacs)
      - [ ] run `renamer -edit` to rename with editor.
  - OUTPUT.
    - [ ] allow a `-action` flag to specify how to apply a name change (profile/editor)
      - [ ] have a `-action=validate` flag which in conjunction with `-profile` checks if
            any file breaks the profile, but doesn't apply the profile.
      - [ ] default to `-action=apply` flag which applies the renaming rules specified in profile.

Features:

  - INPUT.
    - [ ] xxx.
  - RENAMING.
    - [ ] xxx.
  - OUTPUT.
    - [ ] xxx.

Pitfalls:

  - INPUT.
    - [ ] xxx.
  - RENAMING.
    - [ ] test various types circular renames: eg `(a->b,b->a)`
          should not delete anything, but should execute.
    - [ ] test over-naming renames: `(a->c,c->c)`,
          should not delete anything, should not execute.
    - [ ] test colliding renames: `(a->c,b->c)`
          should not delete anything, should not execute.
  - OUTPUT.
    - [ ] xxx.

Done:

  - [X] use this for filesystem mocks: https://github.com/spf13/afero

## ideas

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

Idea: avoid deleting by renaming.

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

