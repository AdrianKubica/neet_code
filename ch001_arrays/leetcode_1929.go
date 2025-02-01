package main

import "fmt"

func main() {
	s := []int{1,2,1}
	s1 := double(s)
	fmt.Println(s1)
}

func double(s []int) []int {
	length := len(s)
	l := make([]int, length*2)
	for i, v := range s {
		l[i], l[i+length] = v,v
	}
	return l
 }