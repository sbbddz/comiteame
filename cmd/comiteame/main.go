package main

import (
  . "github.com/sbbddz/comiteame"
)

func main() {
	commitMessage := GetCommitMessage()
	Info("Commit: " + commitMessage)

  wtree := SetupGitDir()
  AddFiles(wtree)
  CommitFiles(wtree, commitMessage)
}
