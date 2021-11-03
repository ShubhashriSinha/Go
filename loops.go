package main

import "fmt"

func main() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("The value of x is : ", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("The value of i is : ", i)
	// }

	names := []string{"rose", "mummy", "ruchi", "atul", "papa"}

	// for i := 0; i < len(names); i++ {

	// 	fmt.Println("The value of i is : ", names[i], "Hurray")
	// }

	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
}
