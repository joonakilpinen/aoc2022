package main

import (
	"aoc2022/solutions"
	"aoc2022/solutions/day1"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Solve(solver solutions.Solver, input string) string {
	solution := solver.Solve(input)
	log.Printf("Solution:\n%s", solution)
	return solution
}

func Assert(t *testing.T, solver solutions.Solver, input string, expectedResult string) {
	assert.Equal(t, expectedResult, Solve(solver, input))
}

func TestDay1SolveTest(t *testing.T) {
	Assert(t, &day1.Solver{}, GetTestInput(1), "24000")
}

func TestDay1Solve(t *testing.T) {
	Solve(&day1.Solver{}, GetInput(1))
}
