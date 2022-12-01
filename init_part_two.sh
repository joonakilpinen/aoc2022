#!/usr/bin/env bash

mkdir -p "solutions/day${1}part2"

get_assignment() {
  curl --cookie "$(< .cookie)" "https://adventofcode.com/2022/day/${1}" |
  lynx -dump -nolist -stdin |
  sed "0,/--- Part Two/d" |
  sed 's|^|// |'
}

cat <<EOF > "solutions/day${1}part2/solution.go"
package day${1}part2

type Solver struct{}

func (Solver) Solve(input string) string {
	return ""
}

$(get_assignment "${1}")
EOF

cat <<EOF >> aoc2022_test.go

func TestDay${1}Part2SolveTest(t *testing.T) {
	Assert(t, &day${1}part2.Solver{}, GetTestInput(${1}), "")
}

func TestDay${1}Part2Solve(t *testing.T) {
	Solve(&day${1}part2.Solver{}, GetInput(${1}))
}
EOF