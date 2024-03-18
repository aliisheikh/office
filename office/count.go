package main

import "fmt"

func main() {

	var amt int

	fmt.Println("Enter Any Amount: ")
	fmt.Scanf("%d", &amt)
	fmt.Println("You entered:", amt)

	for i := 0; i < amt; i++ {

		if amt >= 500 {
			fh := amt / 500
			amt = amt - (fh * 500)
			fmt.Println("Five Hundred used", fh, "Times")

		} else if amt >= 100 {
			hundred := amt / 100
			amt = amt - (hundred * 100)
			fmt.Println("Hundred Used", hundred, "Times")

		} else if amt >= 50 {
			fifty := amt / 50
			amt = amt - (fifty * 50)

			fmt.Println("Fifty used", fifty, "Times")
		} else if amt >= 20 {
			twenty := amt / 20

			amt = amt - (twenty * 20)

			fmt.Println("Twenty used", twenty, "Times")
		} else if amt >= 10 {
			ten := amt / 10
			amt = amt - (ten * 10)
			fmt.Println("Ten used", ten, "Times")

		} else if amt >= 5 {
			five := amt / 5
			amt = amt - (five * 5)
			fmt.Println("Five used", five, "Times")
		} else if amt >= 2 {
			two := amt / 2
			amt = amt - (two * 2)
			fmt.Println("Two used ", two, "Times")
		} else if amt >= 1 {
			one := amt
			amt = amt - (one * 1)
			fmt.Println("One Used", one, "Times")
		}

	}

}
