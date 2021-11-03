package main

import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "Hello friends, Good morning"
	// fmt.Println(greeting)
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// //fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// //fmt.Println(strings.ToUpper(greeting))
	// //fmt.Println(strings.Index(greeting, "l"))
	// fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 29, 65, 76, 24, 62, 18, 54}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 20)
	fmt.Println(index)

	names := []string{"rose", "atul", "mummy", "papa"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "rose"))
}
