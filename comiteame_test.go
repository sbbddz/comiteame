package comiteame

import (
  "testing"
)

func assertStrings(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got %s, but want %s", got, want)
  }
}

func TestCommitMessage(t *testing.T) {
  t.Run("Commit Format", func(t *testing.T) {

  })
}
