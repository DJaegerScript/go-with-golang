package main

import (
	"fmt"
	"log"
)

func recoveryFunction() {
	if err := recover(); err != nil {
		log.Println("Program Recovered from", err)
	}
}

func panicFunction() {
	defer recoveryFunction()
	fmt.Println("Get Ready!")
	panic("PANIC!!!!")
	fmt.Println("Panic End")
}

func main() {
	fmt.Println("Program Initialized")
	panicFunction()
	fmt.Println("Program end")
}
