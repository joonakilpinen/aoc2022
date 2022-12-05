package day5part1

import (
	"log"
	"strconv"
	"strings"
)

type Solver struct{}

type stack struct {
	number string
	index  int
	crates []string
}

type instruction struct {
	amount      int
	source      int
	destination int
}

func (s *stack) prepend(item ...string) {
	s.crates = append(item, s.crates...)
}

func (s *stack) append(item ...string) {
	s.crates = append(s.crates, item...)
}

func (s *stack) take(amount int) []string {
	var taken []string
	for _, crate := range s.crates[0:amount] {
		taken = append(taken, crate)
	}
	s.crates = s.crates[amount:]
	return taken
}

func reverse(items []string) []string {
	var reversed []string
	for i := len(items) - 1; i >= 0; i-- {
		reversed = append(reversed, items[i])
	}
	return reversed
}

func parseStacks(input []string) []stack {
	indexLine := input[len(input)-1]
	var stacks []stack
	for i, number := range strings.Split(indexLine, "") {
		if number != " " {
			stacks = append(stacks, stack{number, i, []string{}})
		}
	}
	crateLines := input[:len(input)-1]
	for _, line := range crateLines {
		for i := 0; i < len(stacks); i++ {
			s := &stacks[i]
			if len(line) >= s.index {
				char := line[s.index : s.index+1]
				if char != " " {
					s.append(line[s.index : s.index+1])
				}
			}
		}
	}
	return stacks
}

func getInstructionsAndStacks(input string) ([]stack, []string) {
	split := strings.Split(input, "\n\n")
	in := strings.Split(split[0], "\n")
	ins := strings.Split(split[1], "\n")
	return parseStacks(in), ins
}

func parseInstruction(instructionLine string) instruction {
	split := strings.Split(instructionLine, " ")
	amount, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	source, err := strconv.Atoi(split[3])
	if err != nil {
		log.Fatal(err)
	}
	destination, err := strconv.Atoi(split[5])
	if err != nil {
		log.Fatal(err)
	}
	return instruction{amount, source, destination}
}

func rearrangeStacks(stacks []stack, instructions []string) {
	for _, line := range instructions {
		if len(line) > 0 {
			ins := parseInstruction(line)
			source := &stacks[ins.source-1]
			dest := &stacks[ins.destination-1]
			crates := source.take(ins.amount)
			dest.prepend(reverse(crates)...)
		}
	}
}

func (Solver) Solve(input string) string {
	stacks, instructions := getInstructionsAndStacks(input)
	rearrangeStacks(stacks, instructions)
	var output string
	for _, s := range stacks {
		output += s.crates[0]
	}
	return output
}

//
//    The expedition can depart as soon as the final supplies have been
//    unloaded from the ships. Supplies are stored in stacks of marked
//    crates, but because the needed supplies are buried under many other
//    crates, the crates need to be rearranged.
//
//    The ship has a giant cargo crane capable of moving crates between
//    stacks. To ensure none of the crates get crushed or fall over, the
//    crane operator will rearrange them in a series of carefully-planned
//    steps. After the crates are rearranged, the desired crates will be at
//    the top of each stack.
//
//    The Elves don't want to interrupt the crane operator during this
//    delicate procedure, but they forgot to ask her which crate will end up
//    where, and they want to be ready to unload them as soon as possible so
//    they can embark.
//
//    They do, however, have a drawing of the starting stacks of crates and
//    the rearrangement procedure (your puzzle input). For example:
//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
//
// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2
//
//    In this example, there are three stacks of crates. Stack 1 contains two
//    crates: crate Z is on the bottom, and crate N is on top. Stack 2
//    contains three crates; from bottom to top, they are crates M, C, and D.
//    Finally, stack 3 contains a single crate, P.
//
//    Then, the rearrangement procedure is given. In each step of the
//    procedure, a quantity of crates is moved from one stack to a different
//    stack. In the first step of the above rearrangement procedure, one
//    crate is moved from stack 2 to stack 1, resulting in this
//    configuration:
// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
//
//    In the second step, three crates are moved from stack 1 to stack 3.
//    Crates are moved one at a time, so the first crate to be moved (D) ends
//    up below the second and third crates:
//         [Z]
//         [N]
//     [C] [D]
//     [M] [P]
//  1   2   3
//
//    Then, both crates are moved from stack 2 to stack 1. Again, because
//    crates are moved one at a time, crate C ends up below crate M:
//         [Z]
//         [N]
// [M]     [D]
// [C]     [P]
//  1   2   3
//
//    Finally, one crate is moved from stack 1 to stack 2:
//         [Z]
//         [N]
//         [D]
// [C] [M] [P]
//  1   2   3
//
//    The Elves just need to know which crate will end up on top of each
//    stack; in this example, the top crates are C in stack 1, M in stack 2,
//    and Z in stack 3, so you should combine these together and give the
//    Elves the message CMZ.
//
//    After the rearrangement procedure completes, what crate ends up on top
//    of each stack?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
