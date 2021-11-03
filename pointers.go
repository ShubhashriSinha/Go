package main

import "fmt"

func updateMyName(n *string) string {
	*n = "Zeron"
	return *n
}

func main() {

	name1 := "tifa"

	// updateMyName(name1)
	//fmt.Println("Memory address of name is: ", &name1)

	memAdd := &name1
	// fmt.Println("Value at Memory Address: ", memAdd, "is: ", *memAdd)

	fmt.Println(name1)
	updateMyName(memAdd)
	fmt.Println(name1)

}
