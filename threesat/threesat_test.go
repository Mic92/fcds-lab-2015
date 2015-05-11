package threesat

import (
	"testing"
)

func TestSolveable3Sat(t *testing.T) {
	c := []Clause{
		Clause{3, 3, 3},
		Clause{2, 1, -1},
		Clause{-3, -2, -3},
		Clause{2, 1, 2},
	}
	solver := Solver{c, 3}
	solution := solver.Solve()
	if solution == nil {
		t.Fatal("no solution found")
	}
	if *solution != 5 {
		t.Fatalf("expect solution to equal 5, got %d", *solution)
	}
}

func TestUnsolveable3Sat(t *testing.T) {
	c := []Clause{
		Clause{1, 1, 1},
		Clause{-1, -1, -1},
	}
	solver := Solver{c, 1}
	solution := solver.Solve()
	if solution != nil {
		t.Fatal("expect to find no solution, got: %d", *solution)
	}
}
