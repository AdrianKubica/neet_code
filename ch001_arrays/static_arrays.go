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
}

func removeLast(a *[3]int) {
	a[len(a)-1] = 0
}