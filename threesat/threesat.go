package threesat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Clause struct {
	first  int16
	second int16
	third  int16
}

type Solver struct {
	Clauses []Clause
	NVar    uint
}

func (s *Solver) Solve() *int64 {
	if s.NVar == 0 {
		return nil
	}
	iVar := make([]int64, s.NVar)
	for i := uint(0); i < s.NVar; i++ {
		iVar[i] = 1 << i
	}
	maxNumber := iVar[len(iVar)-1] << 1

	nClauses := len(s.Clauses)
	for number := int64(0); number < maxNumber; number++ {
		var round int
		for c, clause := range s.Clauses {
			v1 := clause.first
			v2 := clause.second
			v3 := clause.third

			if (v1 > 0 && (number&iVar[v1-1]) > 0) ||
				(v1 < 0 && (number&iVar[-v1-1]) == 0) ||
				(v2 > 0 && (number&iVar[v2-1]) > 0) ||
				(v2 < 0 && (number&iVar[-v2-1]) == 0) ||
				(v3 > 0 && (number&iVar[v3-1]) > 0) ||
				(v3 < 0 && (number&iVar[-v3-1]) == 0) {
				round = c
				continue // clause is true
			}

			break // clause is false
		}
		if round == (nClauses - 1) {
			return &number
		}
	}
	return nil
}

func New(in *os.File) (*Solver, error) {
	r := bufio.NewReader(in)

	var nClauses, nVar uint
	_, err := fmt.Fscanln(r, &nClauses, &nVar)
	if err != nil {
		return nil, fmt.Errorf("failed to parse header: %v", err)
	}
	solver := Solver{
		make([]Clause, nClauses),
		nVar,
	}
	for i := range solver.Clauses {
		var v1, v2, v3 int16
		_, err := fmt.Fscanln(r, &v1, &v2, &v3)
		if err == io.EOF {
			return nil, fmt.Errorf("not enougth lines found in file: expected %d, got %d", nClauses, i)
		}
		if err != nil {
			return nil, fmt.Errorf("error while reading file %v: %v", in, err)
		}
		solver.Clauses[i] = Clause{v1, v2, v3}
	}
	return &solver, nil
}
