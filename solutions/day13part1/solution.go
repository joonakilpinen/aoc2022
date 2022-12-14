package day13part1

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Solver struct{}

func parsePacketString(line string) (string, []any) {
	var values []any
	for len(line) > 0 {
		if line[0] == '[' {
			rest, content := parsePacketString(line[1:])
			line = rest
			values = append(values, content)
		} else if line[0] == ',' {
			line = line[1:]
		} else if line[0] == ']' {
			return line[1:], values
		} else {
			_, err := strconv.Atoi(line[0:1])
			if err == nil {
				stuffUntilClosingBracket := strings.Split(line, "]")[0]
				stuffUntilComma := strings.Split(stuffUntilClosingBracket, ",")[0]
				num, err := strconv.Atoi(stuffUntilComma)
				if err != nil {
					log.Fatal(err)
				}
				line = line[len(stuffUntilComma):]
				values = append(values, num)
			}
		}
	}
	return "", values
}

type PacketPair struct {
	Left  []any
	Right []any
}

func ParsePacketPairs(input string) []PacketPair {
	var packetPairs []PacketPair
	pairs := strings.Split(input, "\n\n")
	for _, pair := range pairs {
		split := strings.Split(pair, "\n")
		if len(split) < 2 {
			continue
		}
		leftStr := split[0]
		rightStr := split[1]
		_, left := parsePacketString(leftStr[1 : len(leftStr)-1])
		_, right := parsePacketString(rightStr[1 : len(rightStr)-1])
		packetPairs = append(packetPairs, PacketPair{Left: left, Right: right})
	}
	return packetPairs
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func CompareSlices(left any, right any) int {
	leftSlice := left.([]any)
	rightSlice := right.([]any)
	maxLen := max(len(leftSlice), len(rightSlice))
	for i := 0; i < maxLen; i++ {
		if i >= len(leftSlice) {
			return -1
		}
		if i >= len(rightSlice) {
			return 1
		}
		result := ComparePackets(leftSlice[i], rightSlice[i])
		if result != 0 {
			return result
		}
	}
	return 0
}

func ComparePackets(left any, right any) int {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	if leftType.Kind() == reflect.Slice && rightType.Kind() == reflect.Slice {
		return CompareSlices(left, right)
	}
	if leftType.Kind() == reflect.Slice && rightType.Kind() == reflect.Int {
		return CompareSlices(left, []any{right})
	}
	if leftType.Kind() == reflect.Int && rightType.Kind() == reflect.Slice {
		return CompareSlices([]any{left}, right)
	}
	if leftType.Kind() == reflect.Int && rightType.Kind() == reflect.Int {
		return left.(int) - right.(int)
	}
	panic("Some types left unchecked?")
}

func (Solver) Solve(input string) string {
	packetPairs := ParsePacketPairs(input)
	amount := 0
	for i, pair := range packetPairs {
		if ComparePackets(pair.Left, pair.Right) < 0 {
			amount += i + 1
		}
	}
	return strconv.Itoa(amount)
}

//
//    You climb the hill and again try contacting the Elves. However, you
//    instead receive a signal you weren't expecting: a distress signal.
//
//    Your handheld device must still not be working properly; the packets
//    from the distress signal got decoded out of order. You'll need to
//    re-order the list of received packets (your puzzle input) to decode the
//    message.
//
//    Your list consists of pairs of packets; pairs are separated by a blank
//    line. You need to identify how many pairs of packets are in the right
//    order.
//
//    For example:
// [1,1,3,1,1]
// [1,1,5,1,1]
//
// [[1],[2,3,4]]
// [[1],4]
//
// [9]
// [[8,7,6]]
//
// [[4,4],4,4]
// [[4,4],4,4,4]
//
// [7,7,7,7]
// [7,7,7]
//
// []
// [3]
//
// [[[]]]
// [[]]
//
// [1,[2,[3,[4,[5,6,7]]]],8,9]
// [1,[2,[3,[4,[5,6,0]]]],8,9]
//
//    Packet data consists of lists and integers. Each list starts with [,
//    ends with ], and contains zero or more comma-separated values (either
//    integers or other lists). Each packet is always a list and appears on
//    its own line.
//
//    When comparing two values, the first value is called left and the
//    second value is called right. Then:
//      * If both values are integers, the lower integer should come first.
//        If the left integer is lower than the right integer, the inputs are
//        in the right order. If the left integer is higher than the right
//        integer, the inputs are not in the right order. Otherwise, the
//        inputs are the same integer; continue checking the next part of the
//        input.
//      * If both values are lists, compare the first value of each list,
//        then the second value, and so on. If the left list runs out of
//        items first, the inputs are in the right order. If the right list
//        runs out of items first, the inputs are not in the right order. If
//        the lists are the same length and no comparison makes a decision
//        about the order, continue checking the next part of the input.
//      * If exactly one value is an integer, convert the integer to a list
//        which contains that integer as its only value, then retry the
//        comparison. For example, if comparing [0,0,0] and 2, convert the
//        right value to [2] (a list containing 2); the result is then found
//        by instead comparing [0,0,0] and [2].
//
//    Using these rules, you can determine which of the pairs in the example
//    are in the right order:
// == Pair 1 ==
// - Compare [1,1,3,1,1] vs [1,1,5,1,1]
//   - Compare 1 vs 1
//   - Compare 1 vs 1
//   - Compare 3 vs 5
//     - Left side is smaller, so inputs are in the right order
//
// == Pair 2 ==
// - Compare [[1],[2,3,4]] vs [[1],4]
//   - Compare [1] vs [1]
//     - Compare 1 vs 1
//   - Compare [2,3,4] vs 4
//     - Mixed types; convert right to [4] and retry comparison
//     - Compare [2,3,4] vs [4]
//       - Compare 2 vs 4
//         - Left side is smaller, so inputs are in the right order
//
// == Pair 3 ==
// - Compare [9] vs [[8,7,6]]
//   - Compare 9 vs [8,7,6]
//     - Mixed types; convert left to [9] and retry comparison
//     - Compare [9] vs [8,7,6]
//       - Compare 9 vs 8
//         - Right side is smaller, so inputs are not in the right order
//
// == Pair 4 ==
// - Compare [[4,4],4,4] vs [[4,4],4,4,4]
//   - Compare [4,4] vs [4,4]
//     - Compare 4 vs 4
//     - Compare 4 vs 4
//   - Compare 4 vs 4
//   - Compare 4 vs 4
//   - Left side ran out of items, so inputs are in the right order
//
// == Pair 5 ==
// - Compare [7,7,7,7] vs [7,7,7]
//   - Compare 7 vs 7
//   - Compare 7 vs 7
//   - Compare 7 vs 7
//   - Right side ran out of items, so inputs are not in the right order
//
// == Pair 6 ==
// - Compare [] vs [3]
//   - Left side ran out of items, so inputs are in the right order
//
// == Pair 7 ==
// - Compare [[[]]] vs [[]]
//   - Compare [[]] vs []
//     - Right side ran out of items, so inputs are not in the right order
//
// == Pair 8 ==
// - Compare [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
//   - Compare 1 vs 1
//   - Compare [2,[3,[4,[5,6,7]]]] vs [2,[3,[4,[5,6,0]]]]
//     - Compare 2 vs 2
//     - Compare [3,[4,[5,6,7]]] vs [3,[4,[5,6,0]]]
//       - Compare 3 vs 3
//       - Compare [4,[5,6,7]] vs [4,[5,6,0]]
//         - Compare 4 vs 4
//         - Compare [5,6,7] vs [5,6,0]
//           - Compare 5 vs 5
//           - Compare 6 vs 6
//           - Compare 7 vs 0
//             - Right side is smaller, so inputs are not in the right order
//
//    What are the indices of the pairs that are already in the right order?
//    (The first pair has index 1, the second pair has index 2, and so on.)
//    In the above example, the pairs in the right order are 1, 2, 4, and 6;
//    the sum of these indices is 13.
//
//    Determine which pairs of packets are already in the right order. What
//    is the sum of the indices of those pairs?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
