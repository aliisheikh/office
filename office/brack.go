package main

import (
	"fmt"
	"log"
)

func balanced(brackett string) bool {
	// We check each character(brackets) individually.it checks single by single character
	si := []int32{}

	// Now lets looping the value from v to function string
	for _, v := range brackett {
		// strings which we want to compare
		// here we do the check for character by character into trhe string
		if v == '(' || v == '{' || v == '[' {
			// Push the strings into the stack or slice si
			si = append(si, v)

			//fmt.Println(si)

		} else {
			//so, we must pused something into the stack. if the pushed character is not the opening bracket
			// then it must be closing bracket which we do not want to push. so stack are not empty at this point.
			//so we check the length of the slice. if length is 0 means empty return false.

			if len(si) == 0 {
				return false
			}

			// After push character into the slice, it compares the value top of the slice and the new input value which should be same closing bracket
			pop := si[len(si)-1]
			// Pop the value form the slice/stack
			// Simply working as a exclusive slice
			si = si[0 : len(si)-1]

			// Check if the pointed character matches the opening bracket
			//if it is not matched then it returned false
			if (pop == '(' && v != ')') || (pop == '{' && v != '}') || (pop == '[' && v != ']') {
				return false
			}
		}
	}

	// Check if the stack is empty
	return len(si) == 0
}

func main() {

	slice := []string{
		"{}",
		"{]",
		"{)",
		"(}",
		"(]",
		"()",
		"[}",
		"[)",
		"[]",
		"{[(]]}",
		"((())){",
		"{[()]}",
		"((())",
		"[({})]",
		"",
		")(",
		"[{()}]",
		"{[}]",
	}

	for _, v2 := range slice {
		if balanced(v2) {
			fmt.Printf("%s is balanced\n", v2)
		} else {
			fmt.Printf("%s is not balanced\n", v2)
		}
		log.Print(v2)
	}

}
