package day12part2

import (
	"aoc2022/solutions/day12part1"
	"golang.org/x/exp/slices"
	"strconv"
)

type Solver struct{}

func getEligibleNeighbors(heightMap [][]int, visited []day12part1.Coordinates, start day12part1.Coordinates) []day12part1.Coordinates {
	candidates := []day12part1.Coordinates{
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
		{start.X - 1, start.Y},
		{start.X + 1, start.Y},
	}
	var eligibleNeighbors []day12part1.Coordinates
	for _, coord := range candidates {
		if coord.X < 0 {
			continue
		}
		if coord.Y < 0 {
			continue
		}
		if coord.X > len(heightMap)-1 {
			continue
		}
		if coord.Y > len(heightMap[coord.X])-1 {
			continue
		}
		if slices.Contains(visited, coord) {
			continue
		}
		if heightMap[coord.X][coord.Y]+1 < heightMap[start.X][start.Y] {
			continue
		}
		eligibleNeighbors = append(eligibleNeighbors, coord)
	}
	return eligibleNeighbors
}

func findShortestPath(heightMap [][]int, start day12part1.Coordinates, predicate func([][]int, day12part1.Coordinates) bool) []day12part1.Coordinates {
	queue := day12part1.Queue[day12part1.Path]{}
	queue.Add(day12part1.Path{Next: start})
	visited := []day12part1.Coordinates{start}
	for !queue.IsEmpty() {
		path := queue.Poll()
		if predicate(heightMap, path.Next) {
			return path.Route
		}
		for _, neighbor := range getEligibleNeighbors(heightMap, visited, path.Next) {
			visited = append(visited, neighbor)
			queue.Add(day12part1.Path{Next: neighbor, Route: append(path.Route, path.Next)})
		}
	}
	panic("No path?")
}

func (Solver) Solve(input string) string {
	heightMap, _, end := day12part1.ParseHeightMap(input)
	shortestPath := findShortestPath(heightMap, end, func(hm [][]int, c day12part1.Coordinates) bool { return hm[c.X][c.Y] == 1 })
	return strconv.Itoa(len(shortestPath))
}

//
//    As you walk up the hill, you suspect that the Elves will want to turn
//    this into a hiking trail. The beginning isn't very scenic, though;
//    perhaps you can find a better starting point.
//
//    To maximize exercise while hiking, the trail should start as low as
//    possible: elevation a. The goal is still the square marked E. However,
//    the trail should still be direct, taking the fewest steps to reach its
//    goal. So, you'll need to find the shortest path from any square at
//    elevation a to the square marked E.
//
//    Again consider the example from above:
// Sabqponm
// abcryxxl
// accszExk
// acctuvwj
// abdefghi
//
//    Now, there are six choices for starting position (five marked a, plus
//    the square marked S that counts as being at elevation a). If you start
//    at the bottom-left square, you can reach the goal most quickly:
// ...v<<<<
// ...vv<<^
// ...v>E^^
// .>v>>>^^
// >^>>>>>^
//
//    This path reaches the goal in only 29 steps, the fewest possible.
//
//    What is the fewest steps required to move starting from any square with
//    elevation a to the location that should get the best signal?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
