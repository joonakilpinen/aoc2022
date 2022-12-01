#!/usr/bin/env bash

mkdir -p "inputs/day${1}" "solutions/day${1}part1"

get_assignment() {
  curl --cookie "$(< .cookie)" "https://adventofcode.com/2022/day/${1}" |
  lynx -dump -nolist -stdin |
  sed "0,/--- Day ${1}/d" |
  sed 's|^|// |'
}

curl --cookie "$(< .cookie)" "https://adventofcode.com/2022/day/${1}/input" > "inputs/day${1}/input.txt"

cat <<EOF > "solutions/day${1}part1/solution.go"
package day${1}part1

type Solver struct{}

func (Solver) Solve(input string) string {
	return ""
}

$(get_assignment "${1}")
EOF

cat <<EOF >> aoc2022_test.go

func TestDay${1}Part1SolveTest(t *testing.T) {
	Assert(t, &day${1}part1.Solver{}, GetTestInput(${1}), "")
}

func TestDay${1}Part1Solve(t *testing.T) {
	Solve(&day${1}part1.Solver{}, GetInput(${1}))
}
EOF