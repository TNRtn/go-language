package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter input: ")
	fmt.Scanln(&input) // Reads input until a newline (faster than bufio)
	fmt.Println("You entered:", input)
}
