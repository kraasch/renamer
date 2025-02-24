
package main

import (
  // for making a nice centred box.
  tea "github.com/charmbracelet/bubbletea"
  lip "github.com/charmbracelet/lipgloss"

  // argument parser.
  arg "github.com/alexflint/go-arg"

  // basics.
  "fmt"

  // file system.
  "github.com/spf13/afero"
  "os"

  // local packages.
  fsm "github.com/kraasch/renamer/pkg/fsmanage"
)

// TODO: implement the following args:
// NOTE: sketch of the args.
//
// 1) InputMode      (req) = {pipe,dir[+selection],recursive[+selection]}.
// 2) ConversionMode (req) = {rule[+ruleString],profile[+profileName],editor,interactive}.
// 3) OutputMode     (req) = {apply,validate}.
//               selection = {all,files,dirs}.
//              ruleString = STRING
//             profileName = STRING
var args struct {
  Suppress       bool   `arg:"-u"`
  Verbose        bool   `arg:"-v"`
	InputMode      string `arg:"-i,required"`
	ConversionMode string `arg:"-c,required"`
	OutputMode     string `arg:"-o,required"`
  SelectionMode  string `arg:"-s"`
  RuleString     string `arg:"-r"`
  ProfileName    string `arg:"-p"`
}

var (
  // return value.
  output = ""
  // styles.
  styleBox = lip.NewStyle().
     BorderStyle(lip.NormalBorder()).
     BorderForeground(lip.Color("56"))
)

type model struct {
  width  int
  height int
  fs     afero.Fs
}

func (m model) Init() tea.Cmd {
  return func() tea.Msg { return nil }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
  case tea.KeyMsg:
    switch msg.String() {
    case "enter":
      fsm.DirRename(m.fs, "abc.txt", "xyz.txt")
    case "q":
      output = "You quit on me!"
      return m, tea.Quit
    }
  }
  return m, cmd
}

func (m model) View() string {
  var str string
  str += fmt.Sprintf("Verbose:  %#v (%[1]T)\n", args.Verbose)
  str += fmt.Sprintf("Suppress: %#v (%[1]T)", args.Suppress)
  str = styleBox.Render(str)
  return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
}

func main() {

  // use operating system's file system, wrapped in afero.
  fs := afero.NewOsFs()

  // parse flags.
  p := arg.MustParse(&args)
  if args.InputMode != "pipe" && args.InputMode != "dir" && args.InputMode != "recursive" {
    p.Fail("InputMode must be one of {pipe,dir,recursive}")
  }
  if args.ConversionMode != "rule" && args.ConversionMode != "profile" && args.ConversionMode != "editor" && args.ConversionMode != "interactive" {
    p.Fail("ConversionMode must be one of {rule,profile,editor,interactive}")
  }
  if args.OutputMode != "apply" && args.OutputMode != "validate" {
    p.Fail("OutputMode must be one of {apply,validate}")
  }
  if (args.InputMode == "dir" || args.InputMode != "recursive") && args.SelectionMode == "" {
    p.Fail("InputModes {dir,recursive} need a SelectionMode {all,files,dirs}")
  }
  if args.ConversionMode != "rule" && args.RuleString == "" {
    p.Fail("ConversionMode {rule} needs a RuleString argument")
  }
  if args.ConversionMode != "profile" && args.ProfileName == "" {
    p.Fail("ConversionMode {profile} needs a ProfileName argument")
  }

  // init model.
  m := model{0, 0, fs}

  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }

  // print the last highlighted value in calendar to stdout.
  if !args.Suppress {
    fmt.Println(output)
  }

} // fin.

