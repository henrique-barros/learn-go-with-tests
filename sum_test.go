package main

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

func checkSums(got, want []int, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
