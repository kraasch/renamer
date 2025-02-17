
package main

import (
  // for making a nice centred box.
  tea "github.com/charmbracelet/bubbletea"
  lip "github.com/charmbracelet/lipgloss"

  // basics.
  "fmt"
  "flag"

  // file system.
  "github.com/spf13/afero"
  "os"

  // local packages.
  fsm "github.com/kraasch/renamer/pkg/fsmanage"
)

var (
  // return value.
  output = ""
  // flags.
  verbose  = false
  suppress = false
  // styles.
  styleBox = lip.NewStyle().
     BorderStyle(lip.NormalBorder()).
     BorderForeground(lip.Color("56"))
)

type model struct {
  width     int
  height    int
  fs        afero.Fs
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
  if verbose {
    str = "Hello"
  } else {
    str = "Hi"
  }
  str = styleBox.Render(str)
  return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
}

func main() {

  // use operating system's file system, wrapped in afero.
  fs := afero.NewOsFs()

  // parse flags.
  flag.BoolVar(&verbose,  "verbose",   false, "Show info")
  flag.BoolVar(&suppress, "suppress",  false, "Print nothing")
  flag.Parse()

  // init model.
  m := model{0, 0, fs}

  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }

  // print the last highlighted value in calendar to stdout.
  if !suppress {
    fmt.Println(output)
  }

} // fin.

