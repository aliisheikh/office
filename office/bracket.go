package main

import "fmt"

func BalancedBracket(expression string) bool {

	var count1, count2, count3 int

	for _, char := range expression {
		switch char {
		case '(':
			count1++
		case ')':
			count1--
		case '[':
			count2++
		case ']':
			count2--
		case '{':
			count3++
		case '}':
			count3--
		}

		if count1 < 0 || count2 < 0 || count3 < 0 {
			return false
		}
	}

	return count1 == 0 && count2 == 0 && count3 == 0
}

func main() {
	expressions := []string{
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

	for _, exp := range expressions {
		if BalancedBracket(exp) {
			fmt.Printf("%s is balanced\n", exp)
		} else {
			fmt.Printf("%s is not balanced\n", exp)
		}
	}
}
