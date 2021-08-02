package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed at the end")
	fmt.Println("This will be printed at the beginning")
}
