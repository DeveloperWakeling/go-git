package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
  GitFunctions "github.com/developerwakeling/go-git/git"
)

type GitOption int

type Git struct {
  GitOption GitOption
  CommitMessage string

}
// Git option enums
const (
  Fetch GitOption = iota +1
  Pull
  CherryPick
  Commit
  Status
  PullMaster
  GetAllBranches
)

// Convert git option enum in to strings
func (git GitOption) String() string {
  switch git {
  case Fetch:
    return "Fetch"
  case Pull:
    return "Pull"
  case CherryPick:
    return "Cherry Pick"
  case Commit:
    return "Commit"
  case Status:
    return "Status"
  case PullMaster:
    return "Pull Master"
  case GetAllBranches:
    return "Get All Branches"
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
					huh.NewOption("Merge Latest of Master", PullMaster),
					huh.NewOption("Cherry Pick", CherryPick),
					huh.NewOption("Commit", Commit),
					huh.NewOption("Changed Files", Status),
					huh.NewOption("Get All Branches", GetAllBranches),
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
    GitFunctions.Status()
  case Commit:
    GitFunctions.Commit()
  case Fetch:
    GitFunctions.Fetch(true)
  case Pull:
    GitFunctions.Pull(false)
  case PullMaster:
    GitFunctions.Pull(true)
  case GetAllBranches:
    GitFunctions.Fetch(false)
    GitFunctions.GetAllBranches()

  }
}
