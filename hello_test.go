package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", english)
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to the world when no name is supplied", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello in spanish should apply correct prefix in spanish", func(t *testing.T) {
		got := Hello("Enrique", spanish)
		want := "Hola, Enrique"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
