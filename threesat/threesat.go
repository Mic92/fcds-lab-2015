package threesat

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

type Clause struct {
	Value uint64 // the value of the bits
	Mask  uint64 // which bits are set
}

type Solver struct {
	Clauses []Clause
	NVar    uint
}

func (s *Solver) Solve() *uint64 {
	if s.NVar == 0 {
		return nil
	}
	max := uint64(runtime.GOMAXPROCS(-1))
	result := make(chan *uint64, 1)
	maxNumber := uint64(1 << uint(s.NVar))
	for i := uint64(0); i < max; i++ {
		go s.solve(i, max, maxNumber, result)
	}

	for i := uint64(0); i < max; i++ {
		number := <-result
		if number != nil {
			return number
		}
	}
	return nil
}

func (s *Solver) solve(start uint64, step uint64, maxNumber uint64, result chan *uint64) {
	nClauses := len(s.Clauses)
	for number := start; number < maxNumber; number += step {
		var round int
		for c, clause := range s.Clauses {
			// (number XNOR Value) & Mask
			if ((^(number ^ clause.Value)) & clause.Mask) <= 0 {
				break // clause is false
			}
			round = c
		}
		if round == (nClauses - 1) {
			result <- &number
			return
		}
	}

	result <- nil
}

func NewClause(v1, v2, v3 int16) *Clause {
	var value uint64
	uv1 := uint16(-v1)
	uv2 := uint16(-v2)
	uv3 := uint16(-v3)

	if v1 > 0 {
		value |= 1 << uint16(v1-1)
		uv1 = uint16(v1)
	}
	if v2 > 0 {
		value |= 1 << uint16(v2-1)
		uv2 = uint16(v2)
	}
	if v3 > 0 {
		value |= 1 << uint16(v3-1)
		uv3 = uint16(v3)
	}

	if v1 == -v2 || v2 == -v3 || v1 == -v3 {
		return nil // discard -> always true
	}

	return &Clause{
		value,
		1<<(uv1-1) | 1<<(uv2-1) | 1<<(uv3-1),
	}
}

func New(in *os.File) (*Solver, error) {
	r := bufio.NewReader(in)

	var nClauses, nVar uint
	_, err := fmt.Fscanln(r, &nClauses, &nVar)
	if err != nil {
		return nil, fmt.Errorf("failed to parse header: %v", err)
	}
	solver := Solver{
		make([]Clause, 0, nClauses),
		nVar,
	}
	for i := uint(0); i < nClauses; i++ {
		var v1, v2, v3 int16
		_, err := fmt.Fscanln(r, &v1, &v2, &v3)
		if err == io.EOF {
			return nil, fmt.Errorf("not enougth lines found in file: expected %d, got %d", nClauses, i)
		}
		if err != nil {
			return nil, fmt.Errorf("error while reading file %v: %v", in, err)
		}
		c := NewClause(v1, v2, v3)
		if c != nil {
			solver.Clauses = append(solver.Clauses, *c)
		}
	}
	return &solver, nil
}
