package arrays

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		values := []int{1, 3, 4, 5, 6, 7, 8, 9}

		firstEvenNumber := func(val int) bool {
			return val%2 == 0
		}

		found, got := Find(values, firstEvenNumber)
		want := 4

		AssertTrue(t, found)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("find by name", func(t *testing.T) {
		values := []Person{
			{
				Name: "Chris Beck",
			},
			{
				Name: "Kent Fowler",
			},
			{
				Name: "Martin James",
			},
		}

		found, got := Find(values, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
		want := Person{Name: "Chris Beck"}

		AssertTrue(t, found)
		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}

func AssertTrue(t testing.TB, val bool) {
	t.Helper()
	if !val {
		t.Errorf("was expecting value to be %t to be true", val)
	}
}
