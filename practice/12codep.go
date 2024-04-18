package main

import (
	"fmt"
)

// Struct representing a person
type Person struct {
	Name    string
	Age     int
	Address Address // Pointer to Address struct
}

// Struct representing an address
type Address struct {
	Street string
	City   string
}

func main() {
	// Create a new Address struct
	address := Address{Street: "123 Main St", City: "New York"}

	// Create a new Person struct with a pointer to the Address struct
	person := Person{Name: "John", Age: 30, Address: address}

	// Print the details of the person
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Street:", person.Address.Street)
	fmt.Println("City:", person.Address.City)

	// Update the age of the person through the pointer
	person = UpdateAge(person, 45)

	// Print the updated age of the person
	fmt.Println("Updated Age:", person.Age)
}

func UpdateAge(person Person, newAge int) Person {

	person.Age = newAge
	return person
}
