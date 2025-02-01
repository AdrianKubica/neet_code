package main

import (
	"fmt"
	"strings"
)

// READ / WRITE - O(1) - big O of 1
// INSERT / REMOVE END (OVERWRITE) - O(1)
// INSERT MIDDLE - O(N)
// REMOVE MIDDLE - O(N) - big O of N
// memory (capacity) is allocated for all values, even if there are after initialization some default values

func main() {
	a := [5]int{}
	for i := range len(a) {
		a[i] = i
	}
	fmt.Println(a)

	fmt.Println(strings.Repeat("-", 50))
	myArray := [3]int{1,3,5}
	fmt.Println(myArray[1])

	fmt.Println(strings.Repeat("-", 50))
	for i := range myArray {
		fmt.Println(myArray[i])
	}

	i := 0
	for {
		if i < len(myArray) {
			fmt.Println(myArray[i])
			i++
		} else {
			break
		}
	}

	for i :=0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(myArray)
	removeLast(&myArray)
	fmt.Println(myArray)

	fmt.Println(strings.Repeat("-", 50))
	a1 := [...]int{1,2,3}
	removeMiddle(&a1, 1)
	fmt.Println(a1)

	fmt.Println(strings.Repeat("-", 50))
	a2 := [...]int{1,2,3}
	insertMiddle(&a2, 1, 55)
	fmt.Println(a2)
}

func removeLast(a *[3]int) {
	a[len(a)-1] = 0
}

func removeMiddle(a *[3]int, i int) {
	for j := i+1; j < len(a); j++ {
		a[j-1] = a[j]
	}
	a[len(a)-1] = 0
}

func insertMiddle(a *[3]int, idx int, v int) {
	for i := len(a)-1; i > idx; i-- {
		a[i] = a[i-1]
	}
	a[idx]=v
}