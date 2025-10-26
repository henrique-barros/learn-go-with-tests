package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)

		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	checkSums(got, want, t)
}

func TestSumAllTails(t *testing.T) {
	t.Run("correctly sums tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(got, want, t)
	})
	t.Run("correctly calculate for empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0, 2, 9})
		want := []int{0, 11}
		checkSums(got, want, t)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		got := Reduce([]int{1, 3, 4, 5}, func(acc, cur int) int { return acc * cur }, 1)
		want := 60

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("concatenate all elemtns", func(t *testing.T) {
		got := Reduce([]string{"a", "b", "c"}, func(acc, cur string) string { return acc + cur }, "")
		want := "abc"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func checkSums(got, want []int, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
