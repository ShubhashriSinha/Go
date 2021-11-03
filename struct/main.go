package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add Item, s - Save Bill, t - add Tip): ", reader)

	fmt.Println(opt)
}

func main() {

	//var bill1 bill = newBill("bill1")
	// mybill := newBill("mario's bill")
	// mybill.updateTip(20)
	// mybill.addItem("Pie", 120)
	// mybill.addItem("Pizza", 450)
	// mybill.addItem("Paneer Role", 150)

	// fmt.Println(mybill.format())

	myBill1 := createBill()
	promptOptions(myBill1)

	fmt.Println(myBill1)
}
