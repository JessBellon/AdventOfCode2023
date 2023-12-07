package main

type Stack struct {
	s []int
}

func NewStack() *Stack {
	return &Stack{
		s: []int{},
	}
}

func (s *Stack) HasNext() bool {
	return len(s.s) > 0
}

func (s *Stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (int, bool) {
	elements := len(s.s)
	if elements == 0 {
		return 0, false
	}
	n := elements - 1
	v := s.s[n]
	s.s[n] = 0 // Erase element (write zero value)
	s.s = s.s[:n]
	return v, true
}

// func (s Stack) Push(i int) {
// 	s = append(s, i)
// }
//
// func (s Stack) Pop() (int, bool) {
// 	elements := len(s)
// 	if elements == 0 {
// 		return 0, false
// 	}
// 	// Pop
// 	n := elements - 1
// 	val := s[n]
// 	s[n] = 0 // Erase element (write zero value)
// 	s = s[:n]
// 	return val, true
// }
