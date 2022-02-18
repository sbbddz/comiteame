package main

import (
	"fmt"
	"os"
  . "comiteame"
)

func main() {
	commitMessage := GetCommitMessage()
	Info("Commit: " + commitMessage)

	fmt.Println(os.Args[1])
}
