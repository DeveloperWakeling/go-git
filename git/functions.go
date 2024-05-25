package functions

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
)

func Status(){
    cmd := exec.Command("git","status")
    out, err := cmd.Output()
    if err != nil {
      fmt.Println("Error", err)
    }
    fmt.Println(string(out))
}

func Commit(){
    var commitMessage string
    //Start an input
    err := huh.NewInput().
    Title("Commit Message").
    Placeholder("Enter message here").
    Value(&commitMessage).Run()

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
    commitCmd := exec.Command("git", "commit", "-m", commitMessage)
    commitOut, commitError := commitCmd.Output()

    if commitError != nil {
      fmt.Println("commit Error", commitError)
      os.Exit(1)
    }
    fmt.Println(string(commitOut))
}

func Fetch() {
    fetchCmd := exec.Command("git","fetch", "-a")
    out, fetchErr := fetchCmd.Output()
    if fetchErr != nil {
      fmt.Println("Error", fetchErr)
      os.Exit(1)
    }
    fmt.Println(string(out))
}


func Pull(master bool) {
    pullCmd := exec.Command("git","pull")
    if master {
      pullCmd = exec.Command("git","pull", "origin", "master")
    }
    out, pullErr := pullCmd.Output()
    if pullErr != nil {
      fmt.Println("Error", pullErr)
      os.Exit(1)
    }
    fmt.Println(string(out))
}
