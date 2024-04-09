package main

import "fmt"

func main() {

	foo()
	bar("Muhammad ", "Ali")
	s := aloha("ALi")
	fmt.Println(s)

	n, a := dog("Ali", 25)
	fmt.Println(n, a)
}

func foo() {
	fmt.Println("Hi! i'm from function foo. ")

}
func bar(s string, a string) {
	fmt.Println("Hi! My name is ", s, a)

}
func aloha(s string) string {
	return fmt.Sprint("They call me ", s)

}
func dog(name string, age int) (string, int) {
	return fmt.Sprint("Hi My name is ", name, " My age is"), age

}
