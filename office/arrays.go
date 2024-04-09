package main

import "fmt"

func main() {

	fmt.Println("Arrays & Slices")

	//	Arrays

	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := [...]string{"Hello", "Hi", "wohu", "asd", "was that"}
	fmt.Println(y)

	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	//	Hands On Exercise

	/*
		as := [...]string{"Lorem Ipsum is simply", "dummy text of the printing", "and typesetting industry. ", "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."}

			fmt.Println(as)
			fmt.Println(len(as))
			fmt.Printf("%T", as)
	*/

	// Slice Literals

	fmt.Println("Slice Literals\n")

	xs := []string{"Hello", "World", "."}
	fmt.Println(xs)

	//	Hands On Exercise
	hand := []string{"Lorem Ipsum is simply", "dummy text of the printing", "and typesett"}

	fmt.Println(hand)
	fmt.Println(len(hand))

	for i, v := range hand {
		fmt.Printf("index %v - value: %v \n", i, v)

	}

	fmt.Println(hand[0])
	fmt.Println(hand[1])
	fmt.Println(hand[2])

	fmt.Println(len(hand))

	//	Append Function
	fmt.Println("Append Function")
	x2 := []int{12, 23, 34}
	fmt.Println(x2)
	x2 = append(x2, 45, 56, 67, 78)
	fmt.Println("-----------")
	fmt.Println("The Values After Append is:", x2)
	fmt.Println("-----------")

	//	Slicing into slice

	fmt.Println("Slice into slice")

	z := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Slice", z)
	fmt.Println("----------------")

	fmt.Println("EXCLUSIVE SLICE", z[0:9])

	fmt.Println("----------------")

	fmt.Println("Inclusive SLICE", z[1:5])

	//	Deleting Slice

	fmt.Println("Deleting Slice")

	z = append(z[:4], z[5:]...)
	fmt.Println(z)
	fmt.Println("--------------")

	fmt.Println("Make Slice")
	//	Make slice
	si := make([]int, 0, 9)
	fmt.Println(si)
	fmt.Println(len(si))
	fmt.Println(cap(si))
	si = append(si, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(si)

	//	Internal Slice of Array
	a := []int{1, 2, 3, 4, 5, 6}
	b := a
	fmt.Println("The Array of a is:", a)

	fmt.Println("The Value of b is:", b)
	fmt.Println("\n")

	a[0] = 7

	fmt.Println("The Array of a is:", a)
	fmt.Println("The Value of b is:", b)

}
