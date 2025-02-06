package main

import "fmt"

func main() {
	s := Stack[int]{1,2,3,4,5,6}
	fmt.Println(s, len(s), cap(s))
	s.append(101)
	fmt.Println(s, len(s), cap(s))
	x := s.pop()
	fmt.Println(s, len(s), cap(s), x)
	fmt.Println(s.Peek())
}

type Stack[T any] []T

func (s *Stack[T]) append(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

