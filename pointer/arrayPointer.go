package main

import "fmt"

func main() {
	someoneArray := []int{1, 2, 3}
	someoneElseArray := someoneArray

	fmt.Println("Here, someone array's value is", someoneArray)
	fmt.Println("And someone else array's value is", someoneElseArray)

	someoneArray[0] = 10
	fmt.Print("\n")

	fmt.Println("Yeay, someone array's value is changed to", someoneArray)
	fmt.Println("And also someone else array's value is changed to", someoneElseArray)
}
