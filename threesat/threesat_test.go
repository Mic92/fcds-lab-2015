package threesat_test

import (
	sat "github.com/Mic92/fcds-lab-2015/threesat"
	"testing"
)

func TestSolveable3Sat(t *testing.T) {
	c := []sat.Clause{
		*sat.NewClause(3, 3, 3),
		*sat.NewClause(-3, -2, -3),
		*sat.NewClause(2, 1, 2),
	}
	for k, v := range c {
		t.Logf("%d: %v", k, v)
	}
	solver := sat.Solver{c, 3}
	_, solution := solver.Solve()
	if solution == nil {
		t.Fatal("no solution found")
	}
	if *solution != 5 {
		t.Fatalf("expect solution to equal 5, got %d", *solution)
	}
}

func TestAlwaysTrue(t *testing.T) {
	if sat.NewClause(2, 1, -1) != nil {
		t.Fatal("Expect NewClause to return nil for always true expressions")
	}
}

func TestUnsolveable3Sat(t *testing.T) {
	c := []sat.Clause{
		*sat.NewClause(1, 1, 1),
		*sat.NewClause(-1, -1, -1),
	}
	solver := sat.Solver{c, 1}
	_, solution := solver.Solve()
	if solution != nil {
		t.Fatal("expect to find no solution, got: %d", *solution)
	}
}
