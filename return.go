package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	u := strings.ToUpper(n)
	s := strings.Split(u, " ")

	var names []string
	for _, v := range s {
		names = append(names, v[:1])
	}

	if len(names) > 1 {
		return names[0], names[1]
	}

	return names[0], "_"

}

func main() {
	n1, n2 := getInitials("Rose Dalson")
	fmt.Println(n1, n2)

	n3, n4 := getInitials("Atul Singh")
	fmt.Println(n3, n4)

	n5, n6 := getInitials("shiltey")
	fmt.Println(n5, n6)
}
