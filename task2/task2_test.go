package main

import (
	"maps"
	"testing"
)

func TestWordFrequencyCounter(t *testing.T) {
	input := "Golang"
	result := countWordFrequency(input)

	expectedResult := map[string]int{"a": 1, "g": 2, "l": 1, "n": 1, "o": 1}

	if !maps.Equal(expectedResult, result) {
		t.Errorf("Expected: %v, Returned: %v", expectedResult, result)
	}
}

func TestReturnTrueIfIsPalindrome(t *testing.T) {
	input := "NursesRun"
	result := palindromeChecker(input)

	if result != true {
		t.Errorf("Expected: %v, Returned: %v", true, result)
	}
}

func TestReturnFalseIfNotPalindrome(t *testing.T) {
	input := "golang"
	result := palindromeChecker(input)

	if result != false {
		t.Errorf("Expected: %v, Returned: %v", false, result)
	}
}
