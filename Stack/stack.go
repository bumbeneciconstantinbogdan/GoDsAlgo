package dataStructures

// Stack implementation using slices
type Stack[T any] struct {
	data []T
}

// Append elem to stack
func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

// Retrieve last element from stack and a boolean
// that indicate if the stack is  non empty
func (s *Stack[T]) Pop() (T, bool) {
	if s.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	retrieve := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]

	return retrieve, true
}

// Retrieve the current capacity of the stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}
