package main

import "fmt"

func updateName(n string) string {
	n = "Shakira"
	return n
}

func updateMenu(m map[string]float32) {
	m["Coffee"] = 40
}

func main() {

	name := "Tifa"
	name = updateName(name)

	fmt.Println(name)

	menu := map[string]float32{
		"maggi":  20,
		"Coffee": 30,
		"Tea":    20,
	}

	fmt.Println(menu)

	updateMenu(menu)

	fmt.Println(menu)

}
