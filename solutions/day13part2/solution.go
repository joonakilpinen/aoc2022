package day13part2

import (
	"aoc2022/solutions/day13part1"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type Solver struct{}

func dividerIndex(slice [][]any, indexValue int) int {
	for i, item := range slice {
		if len(item) == 0 {
			continue
		}
		if reflect.TypeOf(item[0]).Kind() != reflect.Slice {
			continue
		}
		arr := item[0].([]any)
		if len(arr) == 0 {
			continue
		}
		if reflect.TypeOf(arr[0]).Kind() != reflect.Int {
			continue
		}
		if arr[0].(int) == indexValue {
			return i + 1
		}
	}
	panic(fmt.Sprintf("Value %d could not be found", indexValue))
}

func (Solver) Solve(input string) string {
	packetPairs := day13part1.ParsePacketPairs(input)
	var packetList [][]any
	for _, pair := range packetPairs {
		packetList = append(packetList, pair.Left, pair.Right)
	}
	packetList = append(packetList, []any{[]any{2}}, []any{[]any{6}})
	sort.Slice(packetList, func(i, j int) bool {
		return day13part1.ComparePackets(packetList[i], packetList[j]) < 0
	})
	solution := dividerIndex(packetList, 2) * dividerIndex(packetList, 6)
	return strconv.Itoa(solution)
}

//
//    Now, you just need to put all of the packets in the right order.
//    Disregard the blank lines in your list of received packets.
//
//    The distress signal protocol also requires that you include two
//    additional divider packets:
// [[2]]
// [[6]]
//
//    Using the same rules as before, organize all packets - the ones in your
//    list of received packets as well as the two divider packets - into the
//    correct order.
//
//    For the example above, the result of putting the packets in the correct
//    order is:
// []
// [[]]
// [[[]]]
// [1,1,3,1,1]
// [1,1,5,1,1]
// [[1],[2,3,4]]
// [1,[2,[3,[4,[5,6,0]]]],8,9]
// [1,[2,[3,[4,[5,6,7]]]],8,9]
// [[1],4]
// [[2]]
// [3]
// [[4,4],4,4]
// [[4,4],4,4,4]
// [[6]]
// [7,7,7]
// [7,7,7,7]
// [[8,7,6]]
// [9]
//
//    Afterward, locate the divider packets. To find the decoder key for this
//    distress signal, you need to determine the indices of the two divider
//    packets and multiply them together. (The first packet is at index 1,
//    the second packet is at index 2, and so on.) In this example, the
//    divider packets are 10th and 14th, and so the decoder key is 140.
//
//    Organize all of the packets into the correct order. What is the decoder
//    key for the distress signal?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
