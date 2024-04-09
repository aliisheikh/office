package main

import "fmt"

type soul struct {
	firstname    string
	lastname     string
	fav_icecream []string
}
type vehicle struct {
	engine
	make  string
	model string
	doors string
	color string
}

type engine struct {
	electric bool
}

func main() {
	fmt.Println("\n")
	fmt.Println("Hands on Exercise 53")
	//Hands on Exercise 53
	p1 := soul{
		firstname:    "Muhammad",
		lastname:     "Ali",
		fav_icecream: []string{"Caramel", "Chocolate"},
	}

	p2 := soul{
		firstname:    "Muhammad",
		lastname:     "Umer",
		fav_icecream: []string{"Chocolate", "strawberry"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range p2.fav_icecream {
		fmt.Println(k, v)

	}
	fmt.Println("\n")
	fmt.Println("Hands on Exercise 54")
	m := map[string]soul{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	for _, v := range m {
		fmt.Println(v)

		for _, v2 := range v.fav_icecream {
			fmt.Println(v2)

		}
	}
	fmt.Println("\n")
	fmt.Println("Hands on Exercise 55")

	e1 := vehicle{
		engine: engine{

			electric: true,
		},
		make:  "Civic",
		model: "Rebirth",
		doors: "Electric",
		color: "Grey",
	}
	fmt.Println(e1)

	e2 := vehicle{
		engine: engine{

			electric: true,
		},
		make:  "Honda",
		model: "Civic",
		doors: "Electric",
		color: "Grey",
	}
	fmt.Println(e2)
	fmt.Println("\n")
	//	Hands On Exercise 56
	fmt.Println("Hands On Exercise 56")

	d1 := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "Ahmed",
		friends: map[string]int{
			"Ali":  23,
			"Umer": 21,
			"Qazi": 25,
		},
		favDrink: []string{"Cola", "pepsi", "Fanta"},
	}
	fmt.Println(d1)
	for i, v := range d1.friends {
		fmt.Println(i, v)

	}
}
