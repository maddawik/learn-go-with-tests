package generics

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(value int) {
}

func (s *StackOfInts) Pop() (int, bool) {
	return 0, nil
}

func (s *StackOfInts) IsEmpty() bool {
	return true
}
