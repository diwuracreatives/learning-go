package main

import "strings"

// Word Frequency Count
func countWordFrequency(word string) map[string]int {
	dictionary := map[string]int{}

	for i := range word {
		letter := strings.ToLower(string(word[i]))

		value, ok := dictionary[letter]

		if ok {
			dictionary[letter] = value + 1
		} else {
			dictionary[letter] = 1
		}
	}
	return dictionary
}

// Palindrome Checker
func palindromeChecker(word string) bool {
	length := len(word)
	for i := range word {
		j := length - i - 1
		start := strings.ToLower(string(word[i]))
		end := strings.ToLower(string(word[j]))

		if i <= j {
			if start != end {
				return false
			}
		} else {
			return true
		}
	}
	return true
}
