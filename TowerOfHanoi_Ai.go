// Tower of Hanoi with input validation, legal-move checks, and a final-state
// verification. Run with: go run TowerOfHanoi_Ai.go -disks=4
package main

import (
	"flag"
	"fmt"
	"os"
)

const maxDisks = 20 // Keeps the run practical because the move count is 2^n - 1.

type peg struct {
	name  string
	disks []int // Bottom to top; the smallest disk is at the end.
}

type solver struct {
	moves uint64
	trace bool
}

func (p *peg) moveTopTo(destination *peg) error {
	if len(p.disks) == 0 {
		return fmt.Errorf("cannot move from empty peg %s", p.name)
	}

	disk := p.disks[len(p.disks)-1]
	if len(destination.disks) > 0 && destination.disks[len(destination.disks)-1] < disk {
		return fmt.Errorf("illegal move: disk %d cannot be placed on disk %d", disk, destination.disks[len(destination.disks)-1])
	}

	p.disks = p.disks[:len(p.disks)-1]
	destination.disks = append(destination.disks, disk)
	return nil
}

func (s *solver) solve(n int, source, target, helper *peg) error {
	if n == 0 {
		return nil
	}
	if err := s.solve(n-1, source, helper, target); err != nil {
		return err
	}
	if err := source.moveTopTo(target); err != nil {
		return err
	}
	s.moves++
	if s.trace {
		fmt.Printf("Move %d: disk %d, %s -> %s\n", s.moves, n, source.name, target.name)
	}
	return s.solve(n-1, helper, target, source)
}

func expectedMoves(n int) uint64 {
	return (uint64(1) << uint(n)) - 1
}

func isSolved(n int, source, helper, target *peg) bool {
	if len(source.disks) != 0 || len(helper.disks) != 0 || len(target.disks) != n {
		return false
	}
	for i, disk := range target.disks {
		if disk != n-i {
			return false
		}
	}
	return true
}

func run(n int, trace bool) error {
	source := &peg{name: "A", disks: make([]int, n)}
	for i := range source.disks {
		source.disks[i] = n - i
	}
	helper := &peg{name: "B"}
	target := &peg{name: "C"}
	s := &solver{trace: trace}

	fmt.Printf("Tower of Hanoi: %d disk(s), source %s, target %s, helper %s\n", n, source.name, target.name, helper.name)
	if err := s.solve(n, source, target, helper); err != nil {
		return err
	}

	want := expectedMoves(n)
	if s.moves != want {
		return fmt.Errorf("move count mismatch: got %d, expected %d", s.moves, want)
	}
	if !isSolved(n, source, helper, target) {
		return fmt.Errorf("verification failed: disks did not finish in the target peg")
	}
	fmt.Printf("Solved successfully in %d moves (expected %d).\n", s.moves, want)
	return nil
}

func main() {
	disks := flag.Int("disks", 3, "number of disks (1-20)")
	trace := flag.Bool("trace", true, "print each move; use -trace=false for a summary only")
	flag.Parse()

	if *disks < 1 || *disks > maxDisks {
		fmt.Fprintf(os.Stderr, "invalid disk count %d: choose a value from 1 to %d\n", *disks, maxDisks)
		os.Exit(2)
	}
	if err := run(*disks, *trace); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
