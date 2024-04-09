package main

import "fmt"

func removeDuplicates(dup []int) []int {
	//Empty slice to check the dup values
	a := map[int]bool{}
	//After check, store the unique values in it.
	result := []int{}
	//Check the value to the dup
	for _, v := range dup {
		//Check if the integer is checked or repeat before or not, if it is repeated single time in slice,it will print as it is in slice.
		//if not it will, it will print its unique existence rest will be ignored.
		if !a[v] {
			//it stores the unique value of slice and then return the condition true.
			result = append(result, v)
		}
	}

	return result
}

func main() {

	Before_Dup := []int{1, 1, 2}

	After_Dup := removeDuplicates(Before_Dup)

	fmt.Println(After_Dup)
}
