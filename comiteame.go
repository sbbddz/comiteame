package comiteame

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

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

func SetupGitDir() {

}
