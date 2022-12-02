package day2part2

import (
	"strconv"
	"strings"
)

type rps int

const (
	ROCK     rps = 1
	PAPER    rps = 2
	SCISSORS rps = 3
)

const (
	WIN = iota
	DRAW
	LOSE
)

const winningPrize = 6
const drawPrize = 3

var (
	rpsMap = map[string]rps{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
	}
	strategyMap = map[string]int{
		"X": LOSE,
		"Y": DRAW,
		"Z": WIN,
	}
	winMap = map[rps]rps{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}
	loseMap = map[rps]rps{
		ROCK:     SCISSORS,
		PAPER:    ROCK,
		SCISSORS: PAPER,
	}
)

type Solver struct{}

func iAmTheWinner(opponent rps, me rps) bool {
	return winMap[opponent] == me
}

func getMyMove(opponent rps, strategy int) rps {
	if strategy == DRAW {
		return opponent
	}
	if strategy == LOSE {
		return loseMap[opponent]
	}
	return winMap[opponent]
}

func getLineScore(input string) rps {
	moves := strings.Split(input, " ")
	opponent := rpsMap[moves[0]]
	me := getMyMove(opponent, strategyMap[moves[1]])
	if iAmTheWinner(opponent, me) {
		return me + winningPrize
	}
	if opponent == me {
		return me + drawPrize
	}
	return me
}

func (Solver) Solve(input string) string {
	score := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			score += int(getLineScore(line))
		}
	}
	return strconv.Itoa(score)
}

//
//    The Elf finishes helping with the tent and sneaks back over to you.
//    "Anyway, the second column says how the round needs to end: X means you
//    need to lose, Y means you need to end the round in a draw, and Z means
//    you need to win. Good luck!"
//
//    The total score is still calculated in the same way, but now you need
//    to figure out what shape to choose so the round ends as indicated. The
//    example above now goes like this:
//      * In the first round, your opponent will choose Rock (A), and you
//        need the round to end in a draw (Y), so you also choose Rock. This
//        gives you a score of 1 + 3 = 4.
//      * In the second round, your opponent will choose Paper (B), and you
//        choose Rock so you lose (X) with a score of 1 + 0 = 1.
//      * In the third round, you will defeat your opponent's Scissors with
//        Rock for a score of 1 + 6 = 7.
//
//    Now that you're correctly decrypting the ultra top secret strategy
//    guide, you would get a total score of 12.
//
//    Following the Elf's instructions for the second column, what would your
//    total score be if everything goes exactly according to your strategy
//    guide?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
