package main

import "fmt"

// Tower of Hanoi: move all disks from the source peg to the target peg,
// moving one disk at a time and never placing a larger disk on a smaller one.
// The recursive strategy moves n-1 disks to the helper peg, moves the largest
// disk, then moves the n-1 disks onto it. This takes 2^n - 1 moves.
var moveCount int

func hanoi(n int, source, target, helper string) {
	if n == 0 {
		return
	}
	hanoi(n-1, source, helper, target) // move n-1 disks to helper
	moveCount++
	fmt.Printf("Move disk %d: %s -> %s\n", n, source, target)
	hanoi(n-1, helper, target, source) // move n-1 disks to target
}

func main() {
	disks := 3
	hanoi(disks, "A", "C", "B")
	fmt.Printf("Done! Total moves: %d (expected: %d)\n", moveCount, (1<<disks)-1)
}
