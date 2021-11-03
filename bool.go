package main

import "fmt"

func main() {
	// age := 25
	// fmt.Println(age < 40)
	// fmt.Println(age < 50)

	// if age < 30 {
	// 	fmt.Println("Age is less that 30")
	// } else if age < 40 {
	// 	fmt.Println("Age is less that 40")
	// } else {
	// 	fmt.Println("Age is greater than 40")
	// }

	names := []string{"Rose", "Ruchi", "Atul", "Mummy", "Papa"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue // continue means skip this iteration and start another iteration
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break // breaks out of the loop (Come out of the loop from here)
		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}
}
