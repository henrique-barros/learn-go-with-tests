package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(item T) {
	s.values = append(s.values, item)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, true
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return value, false
}
