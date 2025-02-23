
# !!! UNFINISHED WORK !!!

This program doesn't work and can delete files.

## renamer

## overview

The CLI program is `cmd/renamer.go` and compiles to `renamer`.
The main package is `rnmanage`.
The `pkg` directory below is structured by dependency.

```text
.
├── cmd
│   └── renamer.go // CLI program
└── pkg
    └── rnmanage          // orchestrate renaming (editing, automatic) and file system. 
        ├── edit          // calls editor.
        ├── fsmanage      // deals with file system.
        │   ├── dir       // lists directories.
        │   └── testutil  // creates mock file system.
        └── autorn        // orchestrates renaming.
            ├── profiler  // save renaming rules into profiles.
            └── rename    // renames strings.
```

## tasks

  - INPUT. Main functionality:
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
      - [ ] `renamer -recursive=dirs  -profile music_ogg` as above, but recursive.
    - [ ] allow a way to rename by providing a rule string on CLI.
      - [ ] ie `renamer -rule " (),,-,_,caA,"`
      - [ ] ie `renamer -rule " ();;-;_;caA;" -separator=";"` for a different separator than a comma.

  - RENAMING. Main functionality:
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
    - [ ] implement interactive renaming.
      - [ ] for each file choose method: edit, profile or apply some scripts.

  - RENAMING. Features:
    - [ ] automatically prefix, infix, suffix before file ending:
      - [ ] individual files.
        - [ ] add the current date or file creation date.
          - [ ] yyyy-mm-dd.
          - [ ] yyyy-mm-dd_HH-MM.
          - [ ] precise nano-second-timestamp.
        - [ ] add a random id (decimal, hexadecimal, alphanumerical)
      - [ ] groups of files.
        - [ ] add incrementing id to list of files, also with prefixed zeroes.
        - [ ] rename in different order (`afbecd` to `abcdef`) for scanned
              pages (first front side, then backsides in reverse).

  - RENAMING. Pitfalls:
    - [ ] test various types circular renames: eg `(a->b,b->a)`
          should not delete anything, but should execute.
    - [ ] test over-naming renames: `(a->c,c->c)`,
          should not delete anything, should not execute.
    - [ ] test colliding renames: `(a->c,b->c)`
          should not delete anything, should not execute.

  - OUTPUT. Main functionality:
    - [ ] allow a `-action` flag to specify how to apply a name change (profile/editor)
      - [ ] have a `-action=validate` flag which in conjunction with `-profile` checks if
            any file breaks the profile, but doesn't apply the profile.
      - [ ] default to `-action=apply` flag which applies the renaming rules specified in profile.

  - MISC:
    - [ ] create my own config manager at https://github.com/kraasch/goconf
      - [ ] check if config file exists.
      - [ ] if it doesn't create default config.
      - [ ] if it does exist read it's content to text blob.

Done:

  - [X] use this to parse toml config: https://github.com/BurntSushi/toml/tree/master
  - [X] use this for filesystem mocks: https://github.com/spf13/afero

## notes

  - interesting project, file system mock for `os/exec` : https://github.com/schollii/go-test-mock-exec-command
  - alternative package to parse toml config: https://github.com/pelletier/go-toml

