package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	inputs := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
	}

	for _, s := range inputs {
		fmt.Println("Is a palindrome?", "input:", s, "output:", isPalindrome(s))
	}
}

func isPalindrome(s string) bool {
	lowerString := strings.ToLower(s)

	i := 0
	j := len(lowerString) - 1

	for i < j {
		for !(unicode.IsLetter(rune(lowerString[i])) || unicode.IsDigit(rune(lowerString[i]))) {
			i++
		}

		for !(unicode.IsLetter(rune(lowerString[j])) || unicode.IsDigit(rune(lowerString[j]))) {
			j--
		}

		fmt.Println("compare polyndrome letter", string(lowerString[i]), string(lowerString[j]))

		if lowerString[i] != lowerString[j] {
			return false
		}

		i++
		j--
	}

	return true
}
