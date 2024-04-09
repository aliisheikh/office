package main

import "fmt"

func main() {

	//fmt.Println("Slices")
	arr := [5]int{}
	for i := 0; i < 5; i++ {
		arr[i] = i

	}
	for i, v := range arr {
		fmt.Printf("Index %v and the Type %T and Value %v\n", v, v, i)

	}
	fmt.Println("Hands On Exercise 43")
	//	HandsOnExercise 43
	e := []int{41, 42, 43, 44, 45, 46, 47, 48, 49}
	fmt.Println(e)

	for i, v := range e {
		fmt.Printf("Index %v and the Type %T and Value %v\n", v, v, i)

	}
	fmt.Println("---------------------------")

	fmt.Println("Hands On Exercise 44")

	fmt.Println(e[:6])
	fmt.Println(e[3])
	fmt.Println(e[2:7])
	fmt.Println(e[0:9])

	//	Hands On Exercise

	slice := []int{23, 34, 45}
	fmt.Println(slice)

	slice = append(slice, 56, 67, 78)
	fmt.Println(slice)

	y := []int{12, 32, 43, 54, 60}
	slice = append(slice, y...)
	fmt.Println(slice)

	//	Deleting slice
	fmt.Println("Deleting Slice")

	slice = append(slice[:4], slice[5:]...)
	fmt.Println(slice)

	//	Slice In Slice

	james := []string{"bond", "Shaken", "not stirred"}
	bond := []string{"miss", "Moneypenny", "i am 007"}
	jb := [][]string{james, bond}
	for i, v := range jb {
		fmt.Println(i, v)

		for a, b := range v {
			fmt.Println(a, b)

		}

	}

}
