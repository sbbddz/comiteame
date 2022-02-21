package comiteame

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
)

const DEBUG bool = false;

func SelectTypePrompt() string {
	prompt := promptui.Select{
		Label: "Selecciona tipo de commit",
		Items: []string{"feat", "fix", "chore", "debug",
			"docs", "refactor", "test"},
	}

	_, result, _ := prompt.Run()
	return result
}

func GetCommitMessage() string {
	commitType := SelectTypePrompt()
	prompt := promptui.Prompt{
		Label: "Introduce el mensaje",
	}

	result, _ := prompt.Run()
  return fmt.Sprintf("(%s): %s\n", commitType, result)
}

func SetupGitDir() *git.Worktree {
  currentDirectory, err := os.Getwd()
  CheckError(err)

  read, err := git.PlainOpen(currentDirectory)
  CheckError(err)

  wtree, err := read.Worktree()
  CheckError(err)
  
  if DEBUG {
    status, err := wtree.Status();
    CheckError(err)
    fmt.Println(status)
  }

  Info("¡El repositorio local es correcto!")
  return wtree;
}

func AddFiles(wtree *git.Worktree) {
  prompt := promptui.Select{
    Label: "¿Deseas añadir todos los ficheros modificados?",
    Items: []string{"Sí", "No"},
  }

  _, result, _ := prompt.Run()
  if result != "No" {
    _ , err := wtree.Add(".")
    CheckError(err)

    if DEBUG {
      status, err := wtree.Status();
      CheckError(err)
      fmt.Println(status)
    }
    Info("Añadidos todos los ficheros correctamente.")
  } else {
    Info("Opción no implementada todavía")
  }
}

func CommitFiles(wtree *git.Worktree, commitMessage string) {
  _, err := wtree.Commit(commitMessage, &git.CommitOptions{})
  CheckError(err)

  if DEBUG {
    status, err := wtree.Status();
    CheckError(err)
    fmt.Println(status)
  }
  Info("Commit realizada correctamente")
}
