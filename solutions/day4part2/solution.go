package day4part2

import (
	"log"
	"strconv"
	"strings"
)

type Solver struct{}

type assignment struct {
	start int
	end   int
}

type assignmentPair struct {
	first  assignment
	second assignment
}

func parseAssignment(pair string) assignment {
	split := strings.Split(pair, "-")
	start, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return assignment{start, end}
}

func getAssignmentPairs(lines []string) []assignmentPair {
	var assignmentPairs []assignmentPair
	for _, line := range lines {
		if len(line) > 0 {
			ranges := strings.Split(line, ",")
			assignmentPairs = append(assignmentPairs, assignmentPair{
				parseAssignment(ranges[0]),
				parseAssignment(ranges[1]),
			})
		}
	}
	return assignmentPairs
}

func isBetween(point int, start int, end int) bool {
	return point >= start && point <= end
}

func pointsBetween(first assignment, second assignment) bool {
	return isBetween(first.start, second.start, second.end) || isBetween(first.end, second.start, second.end)
}

func rangesOverlap(pair assignmentPair) bool {
	return pointsBetween(pair.first, pair.second) || pointsBetween(pair.second, pair.first)
}

func rangeOverlapCount(assignmentPairs []assignmentPair) int {
	count := 0
	for _, pair := range assignmentPairs {
		if rangesOverlap(pair) {
			count++
		}
	}
	return count
}

func (Solver) Solve(input string) string {
	lines := strings.Split(input, "\n")
	assignmentPairs := getAssignmentPairs(lines)
	return strconv.Itoa(rangeOverlapCount(assignmentPairs))
}

//
//    It seems like there is still quite a bit of duplicate work planned.
//    Instead, the Elves would like to know the number of pairs that overlap
//    at all.
//
//    In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't
//    overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and
//    2-6,4-8) do overlap:
//      * 5-7,7-9 overlaps in a single section, 7.
//      * 2-8,3-7 overlaps all of the sections 3 through 7.
//      * 6-6,4-6 overlaps in a single section, 6.
//      * 2-6,4-8 overlaps in sections 4, 5, and 6.
//
//    So, in this example, the number of overlapping assignment pairs is 4.
//
//    In how many assignment pairs do the ranges overlap?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
