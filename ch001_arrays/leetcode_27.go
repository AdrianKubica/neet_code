package main

import "fmt"

func main() {
	l := []int{0,1,1,2,2,3,0,4,2}
	x := removeValue(l, 2)
	fmt.Println(l, l[:x], x)
}

func removeValue(a []int, val int) int {
	idx := 0
	for _, v := range a {
		if v != val {
			a[idx] = v
			idx++
		}
	}
	return idx
}