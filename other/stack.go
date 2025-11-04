/*
Implementing the extended version of a Stack, using generic type any.

Implemented Pop, IsEmpty, Top, CopyFromSlice and CopyToSlice

For the unit tests we are doing them manually, not using the test standard library, as this would imply the need for a different file.
*/

package main

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() any {
	pop := s.Top()
	s.vals = s.vals[:len(s.vals)]
	return pop
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) Top() T {
	return s.vals[len(s.vals)-1]
}

func (s *Stack[T]) CopyFromSlice(input []T) {
	s.vals = input
}

func (s *Stack[T]) CopyToSlice() []T {
	output := s.vals[:]
	s.vals = []T{}
	return output
}

func main() {
	TestAll()
}

// Function that panics if any test fails
func TestAll() {
	stack := Stack[int]{}
	stack.vals = []int{1, 2, 3, 4}

	val := stack.Pop()
	if val != 4 && len(stack.vals) != 3 {
		panic("ERROR: stack pop")
	}

	stack.Push(8)
	if stack.Top() != 8 {
		panic("ERROR: stack push")
	}

	if stack.IsEmpty() || !(&Stack[int]{}).IsEmpty() {
		panic("ERROR: IsEmpty")
	}

	if stack.Top() != 8 {
		panic("ERROR Top")
	}

	stack.CopyFromSlice([]int{9})
	if stack.vals[0] != 9 || len(stack.vals) != 1 {
		panic("ERROR CopyFromSlice")
	}

	output := stack.CopyToSlice()
	if !stack.IsEmpty() || output[0] != 9 {
		panic("CopyToSlice")
	}
}
