package main

import "fmt"

// Struct representing a person
type Personn struct {
	Name    string
	Age     int
	Address Address // Not a pointer
}

// Struct representing an address
type Addresses struct {
	Street string
	City   string
}

func main() {
	// Create a new Address struct
	address := Address{Street: "123 Main St", City: "New York"}

	// Create a new Person struct with an Address struct
	person := Person{Name: "John", Age: 30, Address: address}

	// Print the details of the person
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Street:", person.Address.Street)
	fmt.Println("City:", person.Address.City)

	// Update the age of the person without pointers
	person = updateAge(person, 35)

	// Print the updated age of the person
	fmt.Println("Updated Age:", person.Age)
}

// Function to update the age of a person without using pointers
func updateAge(person Person, newAge int) Person {
	// Update the age field
	person.Age = newAge
	return person
}
