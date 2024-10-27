package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk membalikkan satu kata
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Fungsi untuk membalik semua kata dalam kalimat
func reverseEachWord(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = reverseWord(word)
	}
	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat (minimal 3 kata): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if len(strings.Fields(input)) < 3 {
		fmt.Println("Masukkan minimal 3 kata.")
		return
	}

	result := reverseEachWord(input)
	fmt.Println(result)
}
