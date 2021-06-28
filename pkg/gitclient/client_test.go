package gitclient

import (
	"testing"
)

func TestDefaultCommitMsg(t *testing.T) {
	t.Run("It returns a default commit message", func(t *testing.T) {
		got := defaultCommitMsg()
		want := "Merged by gomerge CLI."

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestDefaultApproveMsg(t *testing.T) {
	t.Run("It returns a default approve message.", func(t *testing.T) {
		prId := 1
		got := defaultApproveMsg(prId)
		want := "PR #1 has been approved by [GoMerge](https://github.com/Cian911/gomerge) tool. :rocket:"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
