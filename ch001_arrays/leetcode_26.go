package main

import (
	"fmt"
	"strings"
)

func main() {
	l := []int{0,0,1,1,1,2,2,3,3,4}
	x := removeDuplicated(l)
	fmt.Println(l,l[:x], x)

	fmt.Println(strings.Repeat("-", 50))
	l1 := []int{0,0,1,1,1,2,2,3,3,4}
	x1 := removeDuplicated(l1)
	fmt.Println(l,l[:x1], x1)
}

func removeDuplicated(a []int) int {
	idx := 0
	for _, v := range a {
		if a[idx] != v {
			idx++
			a[idx] = v
		}
	}
	return idx + 1
}

func removeDuplicated2(a []int) int {
	idx := 1
	for _, v := range a {
		if a[idx-1] != v {
			a[idx] = v
			idx++
		}
	}
	return idx
}