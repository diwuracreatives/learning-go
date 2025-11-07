package main

import "testing"

func TestReturnSum(t *testing.T) {
	input := []int{2, 4, 6, 8, 10}
	result := task1(input)

	expectedSum := 30
	if result != expectedSum {
		t.Errorf("Expected: %d, Returned: %d", expectedSum, result)
	}
}

func TestReturnZeroOnEmpty(t *testing.T) {
	input := []int{}
	result := task1(input)

	expectedSum := 0
	if result != expectedSum {
		t.Errorf("Expected: %d, Returned: %d", expectedSum, result)
	}
}
