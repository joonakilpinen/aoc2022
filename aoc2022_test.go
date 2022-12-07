package main

import (
	"aoc2022/solutions"
	"aoc2022/solutions/day1part1"
	"aoc2022/solutions/day1part2"
	"aoc2022/solutions/day2part1"
	"aoc2022/solutions/day2part2"
	"aoc2022/solutions/day3part1"
	"aoc2022/solutions/day3part2"
	"aoc2022/solutions/day4part1"
	"aoc2022/solutions/day4part2"
	"aoc2022/solutions/day5part1"
	"aoc2022/solutions/day5part2"
	"aoc2022/solutions/day6part1"
	"aoc2022/solutions/day6part2"
	"aoc2022/solutions/day7part1"
	"aoc2022/solutions/day7part2"
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

func TestDay3Part1SolveTest(t *testing.T) {
	Assert(t, &day3part1.Solver{}, GetTestInput(3), "157")
}

func TestDay3Part1Solve(t *testing.T) {
	Solve(&day3part1.Solver{}, GetInput(3))
}

func TestDay3Part2SolveTest(t *testing.T) {
	Assert(t, &day3part2.Solver{}, GetTestInput(3), "70")
}

func TestDay3Part2Solve(t *testing.T) {
	Solve(&day3part2.Solver{}, GetInput(3))
}

func TestDay4Part1SolveTest(t *testing.T) {
	Assert(t, &day4part1.Solver{}, GetTestInput(4), "2")
}

func TestDay4Part1Solve(t *testing.T) {
	Solve(&day4part1.Solver{}, GetInput(4))
}

func TestDay4Part2SolveTest(t *testing.T) {
	Assert(t, &day4part2.Solver{}, GetTestInput(4), "4")
}

func TestDay4Part2Solve(t *testing.T) {
	Solve(&day4part2.Solver{}, GetInput(4))
}

func TestDay5Part1SolveTest(t *testing.T) {
	Assert(t, &day5part1.Solver{}, GetTestInput(5), "CMZ")
}

func TestDay5Part1Solve(t *testing.T) {
	Solve(&day5part1.Solver{}, GetInput(5))
}

func TestDay5Part2SolveTest(t *testing.T) {
	Assert(t, &day5part2.Solver{}, GetTestInput(5), "MCD")
}

func TestDay5Part2Solve(t *testing.T) {
	Solve(&day5part2.Solver{}, GetInput(5))
}

func TestDay6Part1SolveTest(t *testing.T) {
	Assert(t, &day6part1.Solver{}, GetTestInput(6), "7")
}

func TestDay6Part1Solve(t *testing.T) {
	Solve(&day6part1.Solver{}, GetInput(6))
}

func TestDay6Part2SolveTest(t *testing.T) {
	Assert(t, &day6part2.Solver{}, GetTestInput(6), "19")
}

func TestDay6Part2Solve(t *testing.T) {
	Solve(&day6part2.Solver{}, GetInput(6))
}

func TestDay7Part1SolveTest(t *testing.T) {
	Assert(t, &day7part1.Solver{}, GetTestInput(7), "95437")
}

func TestDay7Part1Solve(t *testing.T) {
	Solve(&day7part1.Solver{}, GetInput(7))
}

func TestDay7Part2SolveTest(t *testing.T) {
	Assert(t, &day7part2.Solver{}, GetTestInput(7), "24933642")
}

func TestDay7Part2Solve(t *testing.T) {
	Solve(&day7part2.Solver{}, GetInput(7))
}
