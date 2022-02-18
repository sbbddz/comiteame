package main

import (
	"fmt"
	"os"
  . "github.com/sbbddz/comiteame"
)

func main() {
	commitMessage := GetCommitMessage()
	Info("Commit: " + commitMessage)

	fmt.Println(os.Args[1])
}
