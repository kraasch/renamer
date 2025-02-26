
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
  dir "github.com/kraasch/renamer/pkg/dir"
  rnm "github.com/kraasch/renamer/pkg/rnmanage"
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
//              configPath = STRING
var args struct {
  Verbose        bool   `arg:"-v"`
	InputMode      string `arg:"-i,required"`
	ConversionMode string `arg:"-c,required"`
	OutputMode     string `arg:"-o,required"`
  SelectionMode  string `arg:"-s"`
  RuleString     string `arg:"-r"`
  ProfileName    string `arg:"-p"`
	ConfigPath     string `arg:"-C"`
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
  // TODO: implement.
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
  case tea.KeyMsg:
    switch msg.String() {
    case "enter":
      {} // TODO: implement.
    case "q":
      output = "You quit on me!"
      return m, tea.Quit
    }
  }
  return m, cmd
}

func (m model) View() string {
  // TODO: implement.
  var str string
  str += fmt.Sprintf("Verbose:  %#v (%[1]T)\n", args.Verbose)
  str = styleBox.Render(str)
  return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
}

func StartInteractiveGui(fs afero.Fs, in string) string {
  // init model.
  m := model{0, 0, fs}
  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
  return output
}

func StartInteractiveEditor() string {
  return "Have to implement." // TODO: implement.
}

func main() {
  // use operating system's file system, wrapped in afero.
  fs   := afero.NewOsFs()
  iofs := afero.NewIOFS(fs)

  // parse flags.
  p := arg.MustParse(&args)
  // input.
  if args.InputMode != "pipe" && args.InputMode != "dir" && args.InputMode != "recursive" {
    p.Fail("InputMode must be one of {pipe,dir,recursive}")
  }
  if (args.InputMode == "dir" || args.InputMode == "recursive") && args.SelectionMode == "" {
    p.Fail("InputModes {dir,recursive} need a SelectionMode {all,files,dirs}")
  }
  // conversion.
  if args.ConversionMode != "rule" && args.ConversionMode != "profile" && args.ConversionMode != "editor" && args.ConversionMode != "interactive" {
    p.Fail("ConversionMode must be one of {rule,profile,editor,interactive}")
  }
  if args.ConversionMode == "rule" && args.RuleString == "" {
    p.Fail("ConversionMode {rule} needs a RuleString argument")
  }
  if args.ConversionMode == "profile" && args.ProfileName == "" {
    p.Fail("ConversionMode {profile} needs a ProfileName argument")
  }
  if args.ConversionMode != "profile" && args.ConfigPath != "" {
    p.Fail("Only ConversionMode {profile} needs a ConfigPath argument")
  }
  // output.
  if args.OutputMode != "apply" && args.OutputMode != "validate" && args.OutputMode != "print" {
    p.Fail("OutputMode must be one of {apply,validate}")
  }

  // read input.
  input := ""
  switch args.InputMode {
  case "pipe":
    fmt.Println("Get input from pipe")
    input = dir.Pipe()
  case "dir":
    fmt.Println("Get input dir")
    input = dir.DirList(iofs)
  case "recursive":
    fmt.Println("Get input recursively")
    input = dir.DirListTree(iofs)
  }

  // make conversion.
  conversion := ""
  switch args.ConversionMode {
  case "rule":
    fmt.Println("Convert by rule")
    rnm.ConvertByRule(args.RuleString, input)
  case "profile":
    fmt.Println("Convert by profile")
    rnm.AutoConvertByProfile(fs, args.ConfigPath, args.ProfileName, input)
  case "editor":
    fmt.Println("Convert by editor")
    conversion = StartInteractiveEditor()
    rnm.ConvertByPathList(conversion, input)
  case "interactive":
    fmt.Println("Convert interactively")
    conversion = StartInteractiveGui(fs, input)
    rnm.ConvertByPathList(conversion, input)
  }

  // give output.
  switch args.OutputMode {
  case "apply":
    fmt.Println("Give output by applying")
    rnm.ExecuteByApplying()
  case "print":
    fmt.Println("Give output by printing")
    rnm.ExecuteByPrinting()
  case "validate":
    fmt.Println("Give output by validating")
    rnm.ExecuteByValidating()
  }

} // fin.

