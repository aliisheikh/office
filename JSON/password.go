package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginpword1 := `password12309`
	error := bcrypt.CompareHashAndPassword(bs, []byte(loginpword1))
	if error != nil {
		fmt.Println("You Can't LogIn", error)
		return
	}
	fmt.Println("You are logedd In")

}
