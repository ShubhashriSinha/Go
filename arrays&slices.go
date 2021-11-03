package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{2, 3, 4}
	var ages = [3]int{22, 33, 44}

	fmt.Println("ages are : ", ages, " and length is : ", len(ages))

	name := [3]string{"rose", "shub", "rinki"}
	fmt.Println(name)

	//Slices

	var scores = []int{25, 45, 55}
	scores[1] = 50
	scores = append(scores, 80)
	fmt.Println(scores, len(scores))

	// Slice Ranges

	rangeOne := name[0:2]
	rangeTwo := name[1:]
	rangeThree := name[:]

	rangeThree = append(rangeThree, "papa")
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
