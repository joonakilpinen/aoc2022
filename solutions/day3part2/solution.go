package day3part2

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Solver struct{}

func getCommonItem(first string, second string, third string) (rune, error) {
	for _, char := range []rune(first) {
		for _, char2 := range []rune(second) {
			for _, char3 := range []rune(third) {
				if char == char2 && char == char3 {
					return char, nil
				}
			}
		}
	}
	return -1, errors.New("no common value was found")
}

// Couldn't find a better way for this conversion, lol
func getPriority(codepoint rune) rune {
	if unicode.IsUpper(codepoint) {
		return codepoint - 38
	}
	return codepoint - 96
}

func getChunks(size int, lines []string) [][]string {
	var chunks [][]string
	for {
		if len(lines) == 0 {
			break
		}
		if len(lines) < size {
			size = len(lines)
		}
		chunks = append(chunks, lines[0:size])
		lines = lines[size:]
	}
	return chunks
}

func (Solver) Solve(input string) string {
	lines := strings.Split(input, "\n")
	chunks := getChunks(3, lines)
	var sum int32 = 0
	for _, chunk := range chunks {
		if len(chunk) == 3 {
			common, err := getCommonItem(chunk[0], chunk[1], chunk[2])
			if err != nil {
				log.Fatal(err)
			}
			sum += getPriority(common)
		}
	}
	return strconv.FormatInt(int64(sum), 10)
}

//
//    As you finish identifying the misplaced items, the Elves come to you
//    with another issue.
//
//    For safety, the Elves are divided into groups of three. Every Elf
//    carries a badge that identifies their group. For efficiency, within
//    each group of three Elves, the badge is the only item type carried by
//    all three Elves. That is, if a group's badge is item type B, then all
//    three Elves will have item type B somewhere in their rucksack, and at
//    most two of the Elves will be carrying any other item type.
//
//    The problem is that someone forgot to put this year's updated
//    authenticity sticker on the badges. All of the badges need to be pulled
//    out of the rucksacks so the new authenticity stickers can be attached.
//
//    Additionally, nobody wrote down which item type corresponds to each
//    group's badges. The only way to tell which item type is the right one
//    is by finding the one item type that is common between all three Elves
//    in each group.
//
//    Every set of three lines in your list corresponds to a single group,
//    but each group can have a different badge item type. So, in the above
//    example, the first group's rucksacks are the first three lines:
// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
//
//    And the second group's rucksacks are the next three lines:
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw
//
//    In the first group, the only item type that appears in all three
//    rucksacks is lowercase r; this must be their badges. In the second
//    group, their badge item type must be Z.
//
//    Priorities for these items must still be found to organize the sticker
//    attachment efforts: here, they are 18 (r) for the first group and 52
//    (Z) for the second group. The sum of these is 70.
//
//    Find the item type that corresponds to the badges of each three-Elf
//    group. What is the sum of the priorities of those item types?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
