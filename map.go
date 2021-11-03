package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":           100,
		"Tea":            20.5,
		"Pizza":          360.66,
		"Mushroom Salad": 180.9812,
	}

	fmt.Println(menu)
	fmt.Println(menu["Pizza"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		8856822001: "Lucia",
		9725092072: "Jennifer",
		382926328:  "Magro",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[9725092072])

	phonebook[9725092072] = "Luigi"
	phonebook[382926328] = "Yakira"
	fmt.Println(phonebook)
}
