package main

func task1(integers []int) int {

	if len(integers) == 0 {
		return 0
	}

	totalSum := 0
	for i := 0; i < len(integers); i++ {
		totalSum += integers[i]
	}
	return totalSum
}
