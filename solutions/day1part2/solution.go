package day1part2

import (
	"aoc2022/solutions/day1part1"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func (Solver) Solve(input string) string {
	elves := strings.Split(input, "\n\n")
	var cals []int
	for _, elf := range elves {
		cals = append(cals, day1part1.Calories(elf))
	}
	sort.Ints(cals)
	sum := 0
	for i := len(cals) - 1; i > len(cals)-4; i-- {
		sum += cals[i]
	}
	return strconv.Itoa(sum)
}

//
//    By the time you calculate the answer to the Elves' question, they've
//    already realized that the Elf carrying the most Calories of food might
//    eventually run out of snacks.
//
//    To avoid this unacceptable situation, the Elves would instead like to
//    know the total Calories carried by the top three Elves carrying the
//    most Calories. That way, even if one of those Elves runs out of snacks,
//    they still have two backups.
//
//    In the example above, the top three Elves are the fourth Elf (with
//    24000 Calories), then the third Elf (with 11000 Calories), then the
//    fifth Elf (with 10000 Calories). The sum of the Calories carried by
//    these three elves is 45000.
//
//    Find the top three Elves carrying the most Calories. How many Calories
//    are those Elves carrying in total?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
