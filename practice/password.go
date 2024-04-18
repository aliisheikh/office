package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	p := `Password123`
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println("This is yours password ->", p)

	logInPwd := "Password123"

	error := bcrypt.CompareHashAndPassword(b, []byte(logInPwd))
	if error != nil {
		fmt.Println("You can't be Logged In Because Password Doesn't match", error)
		return
	}
	fmt.Println("You are Successfully logged In")

}
