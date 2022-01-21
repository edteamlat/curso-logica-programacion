package main

import (
	"fmt"
	"strings"
)

var numpad = map[int]string{
	0: " ",
	1: "",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func findNumberFromChar(char string) (int, int) {
	for key, value := range numpad {
		index := strings.Index(value, char)
		if index > -1 {
			return key, index + 1
		}
	}

	return -1, 0
}

func printKeys(text string) {
	lastKey := -1
	for _, char := range text {
		key, count := findNumberFromChar(string(char))
		if lastKey == key {
			fmt.Print(" ")
		}
		for count > 0 {
			fmt.Print(key)
			count--
		}
		lastKey = key
	}
	fmt.Println()
}

func main() {
	text := "no"
	printKeys(text)
	text = "si"
	printKeys(text)
	text = "hola mundo"
	printKeys(text)
	text = "alli esta"
	printKeys(text)
}