package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	stack := make([]rune, 0)
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
	for _, v := range s {
        if strings.Contains("([{", string(v)) {
            stack = append(stack, v)
            continue
        }

        if len(stack) == 0 {
            return false
        }   
        
        if val, ok := pairs[v]; ok {
            if stack[len(stack)-1] != val {
                return false
            }
            stack = stack[:len(stack)-1]
        }
	}
	return len(stack) == 0
}