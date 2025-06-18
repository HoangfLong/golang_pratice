package stacks

type Stack struct {
	items []int
}

// push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// peek
func (s *Stack) Peek() int {
	l := len(s.items) - 1
	peek := s.items[l]
	return peek
}

// is empty
// func (s *Stack) isEmpty() int {
// 	isEmpty := true
// 	for()
// }
