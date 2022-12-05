package day5part2

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
			dest.prepend(crates...)
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
//    As you watch the crane operator expertly rearrange the crates, you
//    notice the process isn't following your prediction.
//
//    Some mud was covering the writing on the side of the crane, and you
//    quickly wipe it away. The crane isn't a CrateMover 9000 - it's a
//    CrateMover 9001.
//
//    The CrateMover 9001 is notable for many new and exciting features: air
//    conditioning, leather seats, an extra cup holder, and the ability to
//    pick up and move multiple crates at once.
//
//    Again considering the example above, the crates begin in the same
//    configuration:
//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
//
//    Moving a single crate from stack 2 to stack 1 behaves the same as
//    before:
// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
//
//    However, the action of moving three crates from stack 1 to stack 3
//    means that those three moved crates stay in the same order, resulting
//    in this new configuration:
//         [D]
//         [N]
//     [C] [Z]
//     [M] [P]
//  1   2   3
//
//    Next, as both crates are moved from stack 2 to stack 1, they retain
//    their order as well:
//         [D]
//         [N]
// [C]     [Z]
// [M]     [P]
//  1   2   3
//
//    Finally, a single crate is still moved from stack 1 to stack 2, but now
//    it's crate C that gets moved:
//         [D]
//         [N]
//         [Z]
// [M] [C] [P]
//  1   2   3
//
//    In this example, the CrateMover 9001 has put the crates in a totally
//    different order: MCD.
//
//    Before the rearrangement process finishes, update your simulation so
//    that the Elves know where they should stand to be ready to unload the
//    final supplies. After the rearrangement procedure completes, what crate
//    ends up on top of each stack?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
