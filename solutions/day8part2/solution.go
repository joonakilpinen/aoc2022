package day8part2

import (
	"aoc2022/solutions/day8part1"
	"strconv"
)

type Solver struct{}

func multiply(scores []int) int {
	result := scores[0]
	for _, score := range scores[1:] {
		result *= score
	}
	return result
}

func calculateScenicScore(matrix [][]int, x int, y int) int {
	height := matrix[x][y]
	scores := make([]int, 4)
	for i := x - 1; i >= 0; i-- {
		scores[0]++
		if matrix[i][y] >= height {
			break
		}
	}
	for i := x + 1; i < len(matrix); i++ {
		scores[1]++
		if matrix[i][y] >= height {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		scores[2]++
		if matrix[x][i] >= height {
			break
		}
	}
	for i := y + 1; i < len(matrix[x]); i++ {
		scores[3]++
		if matrix[x][i] >= height {
			break
		}
	}
	return multiply(scores)
}

func calculateScenicScores(matrix [][]int) [][]int {
	scores := make([][]int, len(matrix))
	for x := 0; x < len(matrix); x++ {
		scores[x] = make([]int, len(matrix[x]))
		for y := 0; y < len(matrix[x]); y++ {
			scores[x][y] = calculateScenicScore(matrix, x, y)
		}
	}
	return scores
}

func findHighestScore(matrix [][]int) int {
	highest := 0
	for _, x := range matrix {
		for _, xy := range x {
			if xy > highest {
				highest = xy
			}
		}
	}
	return highest
}

func (Solver) Solve(input string) string {
	matrix := day8part1.ReadMatrix(input)
	scores := calculateScenicScores(matrix)
	highestScore := findHighestScore(scores)
	return strconv.Itoa(highestScore)
}

//
//    Content with the amount of tree cover available, the Elves just need to
//    know the best spot to build their tree house: they would like to be
//    able to see a lot of trees.
//
//    To measure the viewing distance from a given tree, look up, down, left,
//    and right from that tree; stop if you reach an edge or at the first
//    tree that is the same height or taller than the tree under
//    consideration. (If a tree is right on the edge, at least one of its
//    viewing distances will be zero.)
//
//    The Elves don't care about distant trees taller than those found by the
//    rules above; the proposed tree house has large eaves to keep it dry, so
//    they wouldn't be able to see higher than the tree house anyway.
//
//    In the example above, consider the middle 5 in the second row:
// 30373
// 25512
// 65332
// 33549
// 35390
//
//      * Looking up, its view is not blocked; it can see 1 tree (of height
//        3).
//      * Looking left, its view is blocked immediately; it can see only 1
//        tree (of height 5, right next to it).
//      * Looking right, its view is not blocked; it can see 2 trees.
//      * Looking down, its view is blocked eventually; it can see 2 trees
//        (one of height 3, then the tree of height 5 that blocks its view).
//
//    A tree's scenic score is found by multiplying together its viewing
//    distance in each of the four directions. For this tree, this is 4
//    (found by multiplying 1 * 1 * 2 * 2).
//
//    However, you can do even better: consider the tree of height 5 in the
//    middle of the fourth row:
// 30373
// 25512
// 65332
// 33549
// 35390
//
//      * Looking up, its view is blocked at 2 trees (by another tree with a
//        height of 5).
//      * Looking left, its view is not blocked; it can see 2 trees.
//      * Looking down, its view is also not blocked; it can see 1 tree.
//      * Looking right, its view is blocked at 2 trees (by a massive tree of
//        height 9).
//
//    This tree's scenic score is 8 (2 * 2 * 1 * 2); this is the ideal spot
//    for the tree house.
//
//    Consider each tree on your map. What is the highest scenic score
//    possible for any tree?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
