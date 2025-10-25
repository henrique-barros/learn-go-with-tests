package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("string stack", func(t *testing.T) {
		stack := new(Stack[string])

		AssertTrue(t, stack.IsEmpty())

		stack.Push("123")
		AssertFalse(t, stack.IsEmpty())

		stack.Push("456")
		value, _ := stack.Pop()
		AssertEqual(t, value, "456")

		value, _ = stack.Pop()
		AssertEqual(t, value, "123")

		AssertTrue(t, stack.IsEmpty())
	})

	t.Run("int stack", func(t *testing.T) {
		stack := new(Stack[int])

		AssertTrue(t, stack.IsEmpty())

		stack.Push(123)
		AssertFalse(t, stack.IsEmpty())

		stack.Push(456)
		value, _ := stack.Pop()
		AssertEqual(t, value, 456)

		value, _ = stack.Pop()
		AssertEqual(t, value, 123)

		AssertTrue(t, stack.IsEmpty())

		stack.Push(1)
		stack.Push(2)
		firstNum, _ := stack.Pop()
		secondNum, _ := stack.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

func TestAssertFunctions(t *testing.T) {
	t.Run("assert on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 2, 1)
	})

	t.Run("assert on strings", func(t *testing.T) {
		AssertEqual(t, "teste", "teste")
		AssertNotEqual(t, "teste", "teste2")
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
