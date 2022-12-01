# Advent of code 2022

## Prerequisites
- Go
- curl
- lynx

## Instructions

- Login to https://adventofcode.com/
- Copy your session cookie from network tab, and store it in a file named `.cookie`
- Run `./init_day.sh 1` to
  - Create boilerplate for Day 1 tests and solution.
  - Download and save the input for Day 1 in `inputs/day1/input.txt`
  - Write the problem description as comment to `solutions/day1/solution.go`

The init_day.sh -script creates a test function for the test input too,
but this cannot be parsed automatically from the page.
You need to create `inputs/day1/test.txt` manually.

Run tests via [aoc2022_test.go](aoc2022_test.go).
