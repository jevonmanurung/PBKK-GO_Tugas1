package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("input angka: ")
	fmt.Scanln(&input)

	if input == 42 {
		fmt.Println("hello universe")
	} else {
		fmt.Println(input)
	}
}
