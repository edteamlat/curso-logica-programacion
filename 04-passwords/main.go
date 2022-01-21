package main

import (
	"fmt"
	"math/rand"
	"time"
)

var lowercase = "abcdefghijklmnopqrstuvwxyz"
var uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var number = "0123456789"
var special = "!'#$%&^/()@"

func main() {

	rand.Seed(time.Now().UnixNano())
	size := 15
	
	isUpperCase := true
	isNumberCase := true
	isSpecial := true

	fmt.Println(generatePassword(size, isUpperCase, isNumberCase, isSpecial))
}

func generatePassword(size int, isUpperCase, isNumberCase, isSpecial bool) string {
	base := lowercase
	
	if isUpperCase {
		base += uppercase
	}

	if isNumberCase {
		base += number
	}
	
	if isSpecial {
		base += special
	}

	response := ""
	for size > 0 {
		response += string(base[rand.Intn(len(base))])
		size--
	}

	return response
}