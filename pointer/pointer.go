package main

import "fmt"

func main() {
	//* INFO:
	//* change the the type of variable
	//* to either pointer or not

	var someone int = 10
	var someoneElse *int = &someone

	fmt.Println("Hello, someone's value is", someone)
	fmt.Println("And someone else's value is", *someoneElse)

	someone = 20
	fmt.Print("\n")

	fmt.Println("Yeay, someone's value is changed to", someone)
	fmt.Println("and finally, someone else's value is changed to", *someoneElse)

	fmt.Println("Here, someone's memory address is", &someone)
	fmt.Println("And someone else's memory address is", &someoneElse)
}
