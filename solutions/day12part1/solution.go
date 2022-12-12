package day12part1

import (
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type Solver struct{}

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue[T]) Poll() T {
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *Queue[T]) Add(item T) {
	q.queue = append(q.queue, item)
}

type Path struct {
	Next  Coordinates
	Route []Coordinates
}

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) Equals(other Coordinates) bool {
	return c.X == other.X && c.Y == other.Y
}

func getHeight(char rune) int {
	return int(char - 96)
}

func ParseHeightMap(input string) ([][]int, Coordinates, Coordinates) {
	var heightMap [][]int
	var start Coordinates
	var end Coordinates
	for x, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		heightMap = append(heightMap, make([]int, len(line)))
		for y, char := range []rune(line) {
			switch char {
			case 'S':
				start = Coordinates{x, y}
				heightMap[x][y] = getHeight('a')
			case 'E':
				end = Coordinates{x, y}
				heightMap[x][y] = getHeight('z')
			default:
				heightMap[x][y] = getHeight(char)
			}
		}
	}
	return heightMap, start, end
}

func getEligibleNeighbors(heightMap [][]int, visited []Coordinates, start Coordinates) []Coordinates {
	candidates := []Coordinates{
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
		{start.X - 1, start.Y},
		{start.X + 1, start.Y},
	}
	var eligibleNeighbors []Coordinates
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
		if heightMap[coord.X][coord.Y] > heightMap[start.X][start.Y]+1 {
			continue
		}
		eligibleNeighbors = append(eligibleNeighbors, coord)
	}
	return eligibleNeighbors
}

func findShortestPath(heightMap [][]int, start Coordinates, end Coordinates) []Coordinates {
	queue := Queue[Path]{}
	queue.Add(Path{Next: start})
	visited := []Coordinates{start}
	for !queue.IsEmpty() {
		path := queue.Poll()
		if path.Next.Equals(end) {
			return path.Route
		}
		for _, neighbor := range getEligibleNeighbors(heightMap, visited, path.Next) {
			visited = append(visited, neighbor)
			queue.Add(Path{Next: neighbor, Route: append(path.Route, path.Next)})
		}
	}
	panic("No path?")
}

func (Solver) Solve(input string) string {
	heightMap, start, end := ParseHeightMap(input)
	shortestPath := findShortestPath(heightMap, start, end)
	return strconv.Itoa(len(shortestPath))
}

//
//    You try contacting the Elves using your handheld device, but the river
//    you're following must be too low to get a decent signal.
//
//    You ask the device for a heightmap of the surrounding area (your puzzle
//    input). The heightmap shows the local area from above broken into a
//    grid; the elevation of each square of the grid is given by a single
//    lowercase letter, where a is the lowest elevation, b is the
//    next-lowest, and so on up to the highest elevation, z.
//
//    Also included on the heightmap are marks for your current position (S)
//    and the location that should get the best signal (E). Your current
//    position (S) has elevation a, and the location that should get the best
//    signal (E) has elevation z.
//
//    You'd like to reach E, but to save energy, you should do it in as few
//    steps as possible. During each step, you can move exactly one square
//    up, down, left, or right. To avoid needing to get out your climbing
//    gear, the elevation of the destination square can be at most one higher
//    than the elevation of your current square; that is, if your current
//    elevation is m, you could step to elevation n, but not to elevation o.
//    (This also means that the elevation of the destination square can be
//    much lower than the elevation of your current square.)
//
//    For example:
// Sabqponm
// abcryxxl
// accszExk
// acctuvwj
// abdefghi
//
//    Here, you start in the top-left corner; your goal is near the middle.
//    You could start by moving down or right, but eventually you'll need to
//    head toward the e at the bottom. From there, you can spiral around to
//    the goal:
// v..v<<<<
// >v.vv<<^
// .>vv>E^^
// ..v>>>^^
// ..>>>>>^
//
//    In the above diagram, the symbols indicate whether the path exits each
//    square moving up (^), down (v), left (<), or right (>). The location
//    that should get the best signal is still E, and . marks unvisited
//    squares.
//
//    This path reaches the goal in 31 steps, the fewest possible.
//
//    What is the fewest steps required to move from your current position to
//    the location that should get the best signal?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
