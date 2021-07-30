package main

import "fmt"

type someoneStruct struct {
	value int
}

func main() {
	someone := new(someoneStruct)
	someoneElse := someone

	fmt.Println("Here, someone's value is", someone.value)
	fmt.Println("And someone else's value is", someoneElse.value)

	someone.value = 10
	fmt.Print("\n")

	fmt.Println("Yeay, someone's value is changed to", someone.value)
	fmt.Println("And also someone else's value is changed to", someoneElse.value)
}
