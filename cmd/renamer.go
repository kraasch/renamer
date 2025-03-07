
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
  edit "github.com/kraasch/renamer/pkg/edit"
)

// 1) InputMode      (req) = {pipe,dir[+selection],recursive[+selection]}.
// 2) ConversionMode (req) = {rule[+ruleString],profile[+profileName],editor,interactive}.
// 3) OutputMode     (req) = {apply,validate,print}.
//               selection = {all,files,dirs}.
//              ruleString = STRING
//             profileName = STRING
//              configPath = STRING
//              targetDir = STRING
var args struct {
  Verbose        bool   `arg:"-v"`
	InputMode      string `arg:"-i"`
	ConversionMode string `arg:"-c"`
	OutputMode     string `arg:"-o"`
  SelectionMode  string `arg:"-s"`
  RuleString     string `arg:"-r"`
  ProfileName    string `arg:"-p"`
	ConfigPath     string `arg:"-C"`
	TargetDir      string `arg:"-t"`
  ListProfiles   bool   `arg:"-l"`
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

func main() {
  // use operating system's file system, wrapped in afero.
  fs   := afero.NewOsFs()
  iofs := afero.NewIOFS(fs)

  // parse flags.
  p := arg.MustParse(&args)

  // list.
  if args.ListProfiles && ( args.InputMode      != "" || args.ConversionMode != "" || args.OutputMode     != "" || args.SelectionMode  != "" || args.RuleString     != "" || args.ProfileName    != "" || args.ConfigPath     != "" || args.TargetDir      != "") {
    p.Fail("ListProfiles -l must be the only argument.")
  }
  // list profiles.
  if args.ListProfiles {
    fmt.Println(rnm.ListProfiles())
    return
  }

  // input.
  if args.InputMode != "pipe" && args.InputMode != "dir" && args.InputMode != "recursive" {
    p.Fail("InputMode must be one of -i {pipe,dir,recursive}")
  }
  if (args.InputMode == "dir" || args.InputMode == "recursive") && args.SelectionMode == "" {
    p.Fail("InputModes -i {dir,recursive} need a SelectionMode -s {all,files,dirs}")
  }
  // conversion.
  if args.ConversionMode != "rule" && args.ConversionMode != "profile" && args.ConversionMode != "editor" && args.ConversionMode != "interactive" {
    p.Fail("ConversionMode must be one of -c {rule,profile,editor,interactive}")
  }
  if args.ConversionMode == "rule" && args.RuleString == "" {
    p.Fail("ConversionMode -c {rule} needs a RuleString -r 'x|y|z' argument")
  }
  if args.ConversionMode == "profile" && args.ProfileName == "" {
    p.Fail("ConversionMode -c {profile} needs a ProfileName -p 'name' argument")
  }
  if args.ConversionMode != "profile" && args.ConfigPath != "" {
    p.Fail("Only ConversionMode -c {profile} needs a ConfigPath -C 'path' argument")
  }
  // output.
  if args.OutputMode != "apply" && args.OutputMode != "validate" && args.OutputMode != "print" {
    p.Fail("OutputMode must be one of -o {apply,validate,print}")
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
    rnm.ConvertByRule(
      fs,
      args.TargetDir,
      args.RuleString,
      input,
    )
  case "profile":
    fmt.Println("Convert by profile")
    rnm.ConvertByProfile(
      fs,
      args.TargetDir,
      args.ConfigPath,
      args.ProfileName,
      input,
    )
  case "editor":
    fmt.Println("Convert by editor")
    // TODO: also pass in args.TargetDir for more editor awareness.
    conversion = edit.ManualRename(input)
    rnm.ConvertByPathList(
      fs,
      args.TargetDir,
      conversion,
      input,
    )
  case "interactive":
    fmt.Println("Convert interactively")
    conversion = StartInteractiveGui(fs, input)
    rnm.ConvertByPathList(
      fs,
      args.TargetDir,
      conversion,
      input,
    )
  }

  // give output.
  switch args.OutputMode {
  case "apply":
    fmt.Println("Give output by applying")
    rnm.ExecuteByApplying()
  case "print":
    fmt.Println("Give output by formatting")
    fmt.Println(rnm.ExecuteByFormatting())
  case "validate":
    fmt.Print("Give output by validating")
    valid := rnm.ExecuteByValidating()
    if valid {
      fmt.Println("valid.")
    } else {
      fmt.Println("invalid.")
    }
  }

} // fin.

