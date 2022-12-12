package day11part2

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

type Monkey struct {
	items           []int
	operation       func(int) int
	testFunc        func(int) bool
	trueTarget      int
	falseTarget     int
	inspectionCount int
}

func (m *Monkey) hasItems() bool {
	return len(m.items) > 0
}

func (m *Monkey) take() int {
	item := m.items[0]
	m.items = m.items[1:]
	return item
}

func (m *Monkey) put(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) test() bool {
	return m.testFunc(m.items[0])
}

func (m *Monkey) operate(lcm int) {
	m.items[0] = m.operation(m.items[0])
	m.items[0] %= lcm
	m.inspectionCount++
}

func (m *Monkey) getThrowTarget() int {
	if m.test() {
		return m.trueTarget
	}
	return m.falseTarget
}

func (m *Monkey) getInspectedItemAndThrowTarget(lcm int) (int, int) {
	m.operate(lcm)
	target := m.getThrowTarget()
	return m.take(), target
}

func parseStartingItems(line string) []int {
	relevantPart := strings.Split(line, ": ")[1]
	strItems := strings.Split(relevantPart, ", ")
	intItems := make([]int, len(strItems))
	for i, item := range strItems {
		converted, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		intItems[i] = converted
	}
	return intItems
}

func getOpFunction(sign string) func(int, int) int {
	switch sign {
	case "*":
		return func(a int, b int) int {
			return a * b
		}
	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "/":
		return func(a int, b int) int {
			return a / b
		}
	case "-":
		return func(a int, b int) int {
			return a - b
		}
	}
	panic(fmt.Sprintf("Unknown sign: %s", sign))
}

func parseOperation(line string) func(int) int {
	relevantPart := strings.Split(line, "= ")[1]
	strItems := strings.Split(relevantPart, " ")
	opFunc := getOpFunction(strItems[1])
	firstIsOld := strItems[0] == "old"
	lastIsOld := strItems[2] == "old"
	first, firstErr := strconv.Atoi(strItems[0])
	last, lastErr := strconv.Atoi(strItems[2])
	if firstIsOld && lastIsOld {
		return func(old int) int {
			return opFunc(old, old)
		}
	}
	if firstIsOld {
		if lastErr != nil {
			log.Fatal(lastErr)
		}
		return func(old int) int {
			return opFunc(old, last)
		}
	}
	if lastIsOld {
		if firstErr != nil {
			log.Fatal(firstErr)
		}
		return func(old int) int {
			return opFunc(first, old)
		}
	}
	if firstErr != nil || lastErr != nil {
		log.Fatal(firstErr, lastErr)
	}
	return func(old int) int {
		return opFunc(first, last)
	}
}

func parseTest(line string) func(int) bool {
	return func(item int) bool {
		return item%getNumberFromLineEnd(line) == 0
	}
}

func getNumberFromLineEnd(line string) int {
	split := strings.Split(line, " ")
	numStr := split[len(split)-1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseMonkeys(input string) ([]Monkey, int) {
	var monkeys []Monkey
	lcm := 1
	for _, block := range strings.Split(input, "\n\n") {
		lines := strings.Split(block, "\n")
		if len(lines) > 1 {
			monkeys = append(monkeys, Monkey{
				items:           parseStartingItems(lines[1]),
				operation:       parseOperation(lines[2]),
				testFunc:        parseTest(lines[3]),
				trueTarget:      getNumberFromLineEnd(lines[4]),
				falseTarget:     getNumberFromLineEnd(lines[5]),
				inspectionCount: 0,
			})
			lcm *= getNumberFromLineEnd(lines[3])
		}
	}
	return monkeys, lcm
}

func processMonkeys(monkeys []Monkey, lcm int) {
	for i := 0; i < len(monkeys); i++ {
		for monkeys[i].hasItems() {
			item, target := monkeys[i].getInspectedItemAndThrowTarget(lcm)
			monkeys[target].put(item)
		}
	}
}

func (Solver) Solve(input string) string {
	monkeys, lcm := parseMonkeys(input)
	for i := 0; i < 10000; i++ {
		processMonkeys(monkeys, lcm)
	}
	inspectionCounts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspectionCounts[i] = monkey.inspectionCount
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCounts)))
	monkeyBusiness := inspectionCounts[0] * inspectionCounts[1]
	return strconv.Itoa(monkeyBusiness)
}

//
//    You're worried you might not ever get your items back. So worried, in
//    fact, that your relief that a monkey's inspection didn't damage an item
//    no longer causes your worry level to be divided by three.
//
//    Unfortunately, that relief was all that was keeping your worry levels
//    from reaching ridiculous levels. You'll need to find another way to
//    keep your worry levels manageable.
//
//    At this rate, you might be putting up with these monkeys for a very
//    long time - possibly 10000 rounds!
//
//    With these new rules, you can still figure out the monkey business
//    after 10000 rounds. Using the same example above:
// == After round 1 ==
// Monkey 0 inspected items 2 times.
// Monkey 1 inspected items 4 times.
// Monkey 2 inspected items 3 times.
// Monkey 3 inspected items 6 times.
//
// == After round 20 ==
// Monkey 0 inspected items 99 times.
// Monkey 1 inspected items 97 times.
// Monkey 2 inspected items 8 times.
// Monkey 3 inspected items 103 times.
//
// == After round 1000 ==
// Monkey 0 inspected items 5204 times.
// Monkey 1 inspected items 4792 times.
// Monkey 2 inspected items 199 times.
// Monkey 3 inspected items 5192 times.
//
// == After round 2000 ==
// Monkey 0 inspected items 10419 times.
// Monkey 1 inspected items 9577 times.
// Monkey 2 inspected items 392 times.
// Monkey 3 inspected items 10391 times.
//
// == After round 3000 ==
// Monkey 0 inspected items 15638 times.
// Monkey 1 inspected items 14358 times.
// Monkey 2 inspected items 587 times.
// Monkey 3 inspected items 15593 times.
//
// == After round 4000 ==
// Monkey 0 inspected items 20858 times.
// Monkey 1 inspected items 19138 times.
// Monkey 2 inspected items 780 times.
// Monkey 3 inspected items 20797 times.
//
// == After round 5000 ==
// Monkey 0 inspected items 26075 times.
// Monkey 1 inspected items 23921 times.
// Monkey 2 inspected items 974 times.
// Monkey 3 inspected items 26000 times.
//
// == After round 6000 ==
// Monkey 0 inspected items 31294 times.
// Monkey 1 inspected items 28702 times.
// Monkey 2 inspected items 1165 times.
// Monkey 3 inspected items 31204 times.
//
// == After round 7000 ==
// Monkey 0 inspected items 36508 times.
// Monkey 1 inspected items 33488 times.
// Monkey 2 inspected items 1360 times.
// Monkey 3 inspected items 36400 times.
//
// == After round 8000 ==
// Monkey 0 inspected items 41728 times.
// Monkey 1 inspected items 38268 times.
// Monkey 2 inspected items 1553 times.
// Monkey 3 inspected items 41606 times.
//
// == After round 9000 ==
// Monkey 0 inspected items 46945 times.
// Monkey 1 inspected items 43051 times.
// Monkey 2 inspected items 1746 times.
// Monkey 3 inspected items 46807 times.
//
// == After round 10000 ==
// Monkey 0 inspected items 52166 times.
// Monkey 1 inspected items 47830 times.
// Monkey 2 inspected items 1938 times.
// Monkey 3 inspected items 52013 times.
//
//    After 10000 rounds, the two most active monkeys inspected items 52166
//    and 52013 times. Multiplying these together, the level of monkey
//    business in this situation is now 2713310158.
//
//    Worry levels are no longer divided by three after each item is
//    inspected; you'll need to find another way to keep your worry levels
//    manageable. Starting again from the initial state in your puzzle input,
//    what is the level of monkey business after 10000 rounds?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
