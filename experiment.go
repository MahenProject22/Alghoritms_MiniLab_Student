package main

import "fmt"

// countMoves returns how many moves the recursive Tower of Hanoi solution
// needs for n disks. It counts moves without printing every disk movement.
func countMoves(n int) int {
	if n <= 0 {
		return 0
	}
	return 2*countMoves(n-1) + 1
}

func main() {
	fmt.Println("Tower of Hanoi Experiment")
	fmt.Println("Comparing the recursive move count for different numbers of disks:")
	fmt.Printf("%-8s %-14s %-14s\n", "Disks", "Moves counted", "Expected moves")

	for disks := 1; disks <= 6; disks++ {
		actual := countMoves(disks)
		expected := (1 << disks) - 1
		fmt.Printf("%-8d %-14d %-14d\n", disks, actual, expected)
	}

	fmt.Println("\nObservation: each additional disk doubles the previous move count and adds one.")
	fmt.Println("The total number of moves follows the formula 2^n - 1, where n is the number of disks.")
}
