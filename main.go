package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
)

type GitOption int

type Git struct {
  GitOption GitOption
  CommitMessage string

}
// Git option enums
const (
  Fetch GitOption = iota +1
  Merge
  Pull
  CherryPick
  Commit
  Status
)

// Convert git option enum in to strings
func (git GitOption) String() string {
  switch git {
  case Fetch:
    return "Fetch"
  case Merge:
    return "Merge"
  case Pull:
    return "Pull"
  case CherryPick:
    return "Cherry Pick"
  case Commit:
    return "Commit"
  case Status:
    return "Status"
  default:
    return ""
  }
}

func main(){
  var gitOption GitOption
  git:= Git{ GitOption: gitOption}

  form := huh.NewForm(
    huh.NewGroup(
      huh.NewNote().
      Title("Git Branch Manager").
      Description("Manage your branches\n\n").
      Next(true),
    ),
    huh.NewGroup(
      huh.NewSelect[GitOption]().
        Title("Options").
        Options(
          huh.NewOption("Fetch Latest Branches", Fetch).Selected(true),
					huh.NewOption("Pull Latest of Branch", Pull),
					huh.NewOption("Merge", Merge),
					huh.NewOption("Cherry Pick", CherryPick),
					huh.NewOption("Commit", Commit),
					huh.NewOption("Changed Files", Status),
        ).
        Value(&git.GitOption),
    ),
  )  

  err := form.Run()

  if err != nil {
    fmt.Println("Error", err)
    os.Exit(1)
  }

  switch git.GitOption {
  case Status:
    cmd := exec.Command("git","status")
    out, err := cmd.Output()
    if err != nil {
      fmt.Println("Error", err)
    }
    fmt.Println(string(out))
  case Commit:
    //Start an input
    commitForm := huh.NewInput().
    Title("Commit Message").
    Placeholder("Enter message here").
    Value(&git.CommitMessage)

    err := commitForm.Run()

    if err != nil {
      fmt.Println("Error", err)
      os.Exit(1)
    }


    addCmd := exec.Command("git","add", ".")
    _, addErr := addCmd.Output()
    if addErr != nil {
      fmt.Println("Error", addErr)
      os.Exit(1)
    }
      fmt.Println("Commit Message", git.CommitMessage)
    commitCmd := exec.Command("git", "commit", "-m", "'"+git.CommitMessage+"'")
    commitOut, commitError := commitCmd.Output()

    if commitError != nil {
      fmt.Println("commit Error", commitError)
      os.Exit(1)
    }
    fmt.Println(string(commitOut))


  }
}
