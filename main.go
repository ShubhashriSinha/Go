package main

import "fmt"

var age int = 21
var name string = "shub"

func main() {

	// fmt.Println("Hello, Rose")

	// var nameOne string = "mario"
	// var nameTwo = "pokemon"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "mickey"
	// nameThree = "Donald"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "Doraemon"

	// fmt.Println(nameFour)

	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 50

	// fmt.Println(ageOne, ageTwo, ageThree)

	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	//    Println

	fmt.Println("Hello Ninjas!")
	fmt.Println("good bye Ninjas")
	fmt.Println("my age is", age, "and my name is", name)

	// printf (formatted strings) %_ = format specifier

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("My age is of type %T \n", age)
	fmt.Printf("You scored %0.1f points \n", 224.55)

	// Sprintf

	var str = fmt.Sprintf("Hi, my age is %v and my name is %v \n", age, name)
	fmt.Printf("The saved String is : %v", str)
}
