package main

import (
	"aoc2022/solutions"
	"aoc2022/solutions/day1part1"
	"aoc2022/solutions/day1part2"
	"aoc2022/solutions/day2part1"
	"aoc2022/solutions/day2part2"
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

func TestDay1Part1SolveTest(t *testing.T) {
	Assert(t, &day1part1.Solver{}, GetTestInput(1), "24000")
}

func TestDay1Part1Solve(t *testing.T) {
	Solve(&day1part1.Solver{}, GetInput(1))
}

func TestDay1Part2SolveTest(t *testing.T) {
	Assert(t, &day1part2.Solver{}, GetTestInput(1), "45000")
}

func TestDay1Part2Solve(t *testing.T) {
	Solve(&day1part2.Solver{}, GetInput(1))
}

func TestDay2Part1SolveTest(t *testing.T) {
	Assert(t, &day2part1.Solver{}, GetTestInput(2), "15")
}

func TestDay2Part1Solve(t *testing.T) {
	Solve(&day2part1.Solver{}, GetInput(2))
}

func TestDay2Part2SolveTest(t *testing.T) {
	Assert(t, &day2part2.Solver{}, GetTestInput(2), "12")
}

func TestDay2Part2Solve(t *testing.T) {
	Solve(&day2part2.Solver{}, GetInput(2))
}
