package main

import "fmt"

func main() {

	fmt.Println("Maps HandsOn Exercise 49")

	//Maps HandsOn Exercise
	a := make(map[string][]string)
	a["Bond_James"] = []string{"shaken, not Stirred", "martins", "fast cars"}
	a["Money-Penny_Jenny"] = []string{"intelligence", "literature", "computer science"}
	a["No_Dr"] = []string{"cats", "Ice-cream", "sunset"}

	//fmt.Println(a, "\n")

	for k, v := range a {
		fmt.Println(k, v, "\n")

	}

	//	HandsOn Exercise

	fmt.Println("Maps HandsOn Exercise 50")

	//Maps HandsOn Exercise(Add Row in the code)
	b := make(map[string][]string)
	b["Bond_James"] = []string{"shaken, not Stirred", "martins", "fast cars"}
	b["Money-Penny_Jenny"] = []string{"intelligence", "literature", "computer science"}
	b["No_Dr"] = []string{"cats", "Ice-cream", "sunset"}
	b["flemin-ion"] = []string{"steaks", "cigars", "espresso"}

	//fmt.Println(a, "\n")

	for k, v := range b {
		fmt.Println(k, v, "\n")
	}
	fmt.Println("\n")

	fmt.Println("Hands On Exercise 51 ")

	fmt.Println("Delete the record")
	delete(b, "flemin-ion")
	for k, v := range b {
		fmt.Println(k, v, "\n")
	}

}
