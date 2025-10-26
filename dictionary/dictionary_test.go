package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		dictionary.Add("novo", "new def")
		assertDefinition(t, dictionary, "novo", "new def")
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "new def")
		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		newValue := "this is another test"
		dictionary.Update("test", newValue)

		assertDefinition(t, dictionary, "test", newValue)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		newValue := "this is another test"
		err := dictionary.Update("new_test", newValue)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this just an text"}
		err := dictionary.Delete("test")
		assertNoError(t, err)

		_, err = dictionary.Search("test")
		assertError(t, err, ErrNotFound)
	})
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this just an text"}
		err := dictionary.Delete("asd")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s, given %s", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, term string, value string) {
	t.Helper()
	got, err := dictionary.Search(term)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, value)
}
