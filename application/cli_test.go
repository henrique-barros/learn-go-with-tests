package application_test

import (
	"strings"
	"testing"

	app "example.com/hello/application"
)

func TestCLI(t *testing.T) {
	t.Run("record win", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := app.NewStubPlayerStore()
		cli := app.NewCLI(&playerStore, in)

		cli.PlayPoker()

		app.AssertPlayerWin(t, &playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := app.NewStubPlayerStore()

		cli := app.NewCLI(&playerStore, in)
		cli.PlayPoker()

		app.AssertPlayerWin(t, &playerStore, "Cleo")
	})
}
