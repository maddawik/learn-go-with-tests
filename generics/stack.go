package generics

type (
	StackOfStrings = Stack
	StackOfInts    = Stack
)

type Stack struct {
	values []any
}

func (s *Stack) Push(value any) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	index := len(s.values) - 1
	result := s.values[index]
	s.values = s.values[:index]

	return result, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}
