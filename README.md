
# renamer

## demo

<p align="center">
  <img src="./resources/demo.gif" />
</p>

## features

Rename files with renaming rule string (search files)

```bash
renamer -i recursive -s all -o apply -c rule -r " ()||-|_|cAa|"
```

Rename files with renaming rule string (from pipe)

```bash
find | renamer -i pipe -o apply -c rule -r " ()||-|_|cAa|"
```

Edit files with your editor (for example vim or emacs)

```bash
renamer -i recursive -s all -o apply -c editor
```

Different ways to use `renamer`:

  - input:
    - [X] dir: list directory (like `ls`)
    - [X] recursive: search directory tree (like `find`)
    - [X] pipe: pipe in the input.
  - input:
    - [X] rule: provide a rule string.
    - [X] editor: edit the input yourself.
    - [X] config+profile: provide a configuration file with renaming profiles.
    - [ ] interactive: apply rules and profiles with a TUI application.
  - output:
    - [X] apply: rename according to given pattern.
    - [X] print: print what would be done, if apply option was given.
    - [X] validate: print if file system conforms to given rules (i.e. no renaming needed.)

The profile and the rule format are equally powerful, i.e. can be converted into each other.
For now see the [rename](https://github.com/kraasch/renamer/blob/main/pkg/rename/main_test.go) package to see features of rules and profiles.

Example configuration file format:

```toml
# An Example Config
title = "Some Example"

[profiles]
[profiles.toast-txt]
    name            = "n0"
    [profiles.toast-txt.profile_rule]
    word_separators = " ()"
    delete_chars    = ""
    small_gap_mark  = "-"
    big_gap_mark    = "_"
    conversions     = "caA"
    modes_string    = ""
[profiles.prettify-txt]
    name            = "n1"
    [profiles.prettify-txt.profile_rule]
    word_separators = " ()"
    delete_chars    = ""
    small_gap_mark  = "-"
    big_gap_mark    = "_"
    conversions     = "cAa"
    modes_string    = ""
```

## tip

Set an alias for your common renames in your `~/.bashrc`.

```bash
alias rn_mp3='renamer -i recursive -s all -o apply -c rule -r " ()||-|_|cAa|"'
alias rn_txt='renamer -i recursive -s all -o apply -c rule -r " ()||-|_|caA|"'
alias rename='renamer -i dir       -s all -o apply -c editor'
```

This is an alternative to setting up a config file with profiles.

## overview

The CLI program is `cmd/renamer.go` and compiles to `renamer`.
The main package is `rnmanage`.
In the below tree, the `pkg` tree is loosely representing the packages' dependencies.

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

  - renaming: main functionality.
    - [ ] implement interactive renaming.
      - [ ] for each file choose method: edit, profile or apply some scripts.
  - renaming: features.
    - [ ] automatically prefix, infix, suffix before file ending:
      - [ ] individual files.
        - [ ] add the current date.
          - [ ] yyyy-mm-dd.
          - [ ] yyyy-mm-dd_HH-MM.
          - [ ] precise nano-second-timestamp.
        - [ ] add the file creation date.
          - [ ] yyyy-mm-dd.
          - [ ] yyyy-mm-dd_HH-MM.
          - [ ] precise nano-second-timestamp.
        - [ ] add a random id (decimal, hexadecimal, alphanumerical)
    - [ ] groups of files.
      - [ ] add incrementing id to list of files, also with prefixed zeroes.
    - [ ] rename scanned pages.
      - [ ] rename in different order (`afbecd` to `abcdef`) for scanned pages (first front side, then backsides in reverse).
      - [ ] example sorting for 6 PNG files: ABCDEF
      - [ ] split the files in first and second half: --> ABC and DEF
      - [ ] reverse order of second half: --> ABC and FED
      - [ ] alternatingly merge the two halfs back together: --> AFBECD.
  - renaming: pitfalls.
    - [ ] test various types circular renames: eg `(a->b,b->a)` should not delete anything, but should execute.
    - [ ] test over-naming renames: `(a->c,c->c)`, should not delete anything, should not execute.
    - [ ] test colliding renames: `(a->c,b->c)` should not delete anything, should not execute.
  - misc.
    - [ ] also pull out `testutil` package.

Maybe:

  - [ ] Allow separator for rule strings `renamer -rule " ();;-;_;caA;" -separator=";" ...` for a different separator than a pipe.

Done:

  - misc.
    - [X] use this to parse toml config: https://github.com/BurntSushi/toml/tree/master
    - [X] use this for filesystem mocks: https://github.com/spf13/afero
    - [X] create my own config manager at https://github.com/kraasch/goconf
      - [X] check if config file exists.
      - [X] if it doesn't create default config.
      - [X] if it does exist read it's content to text blob.
  - input: main functionality.
    - [X] get file list from pipe, ie allow these inputs:
      - [X] `ls | grep -E 'mp3$' | renamer -edit`
      - [X] `find | grep -E '.ogg$' | renamer -profile music_ogg`
    - [X] if nothing is piped in, have two search modes: `recusvie` and `dir`
      - [X] `renamer -i dir       -s all   ...` list all in pwd (default).
      - [X] `renamer -i dir       -s files ...` list all files in pwd.
      - [X] `renamer -i dir       -s dirs  ...` list all directories in pwd.
      - [X] `renamer -i recursive -s all   ...` as above, but recursive (default)
      - [X] `renamer -i recursive -s files ...` as above, but recursive.
      - [X] `renamer -i recursive -s dirs  ...` as above, but recursive.
      - [X] `renamer -i recursive -s dirs  ...` as above, but recursive.
    - [X] allow a way to rename by providing a rule string on CLI.
      - [X] ie `renamer -rule " ()||-|_|caA| ..."`
  - output: main functionality.
    - [X] allow a `-action` flag to specify how to apply a name change (profile/editor)
      - [X] have a `-action=validate` flag which in conjunction with `-profile` checks if any file breaks the profile, but doesn't apply the profile.
      - [X] default to `-action=apply` flag which applies the renaming rules specified in profile.
  - renaming: main functionality.
    - [X] implement renaming profiles.
      - [X] run `renamer -profile media` to rename media files with specified
            conversion (named `media`, for renaming music, video, etc).
      - [X] default to `renamer -profile default` if no profile is specified.
    - [X] implement pure editor version (use native editor, eg vi, emacs)
      - [X] run `renamer -edit` to rename with editor.

## word of caution

This is work in progress.
This program is

  - powerful: interprets little input to do big things,
  - risky: might delete by renaming,
  - not tested enough yet.

## notes

  - interesting project, file system mock for `os/exec` : https://github.com/schollii/go-test-mock-exec-command
  - alternative package to parse toml config: https://github.com/pelletier/go-toml

