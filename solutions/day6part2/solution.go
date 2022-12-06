package day6part2

import "strconv"

type Solver struct{}

func lastNAreUnique(input string, n int, currentIndex int) bool {
	set := map[uint8]bool{}
	for i := currentIndex; i >= currentIndex-(n-1); i-- {
		set[input[i]] = true
	}
	return len(set) == n
}

func (Solver) Solve(input string) string {
	n := 14
	for i := n - 1; i < len(input); i++ {
		if lastNAreUnique(input, n, i) {
			// Indexing on assignment starts from 1
			return strconv.Itoa(i + 1)
		}
	}
	panic("Something strange in the neighborhood")
}

//
//    Your device's communication system is correctly detecting packets, but
//    still isn't working. It looks like it also needs to look for messages.
//
//    A start-of-message marker is just like a start-of-packet marker, except
//    it consists of 14 distinct characters rather than 4.
//
//    Here are the first positions of start-of-message markers for all of the
//    above examples:
//      * mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
//      * bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
//      * nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
//      * nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
//      * zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26
//
//    How many characters need to be processed before the first
//    start-of-message marker is detected?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
