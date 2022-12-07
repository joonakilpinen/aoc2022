package day7part2

import (
	"aoc2022/solutions/day7part1"
	"sort"
	"strconv"
)

type Solver struct{}

func (Solver) Solve(input string) string {
	tree := day7part1.ParseDirectoryTree(input)
	sizeDirPairs := day7part1.GetTotalSizes(tree)
	sort.Sort(day7part1.SizeDirPairList(sizeDirPairs))
	totalSpace := 70000000
	neededSpace := 30000000
	currentSpace := totalSpace - tree.GetTotalSize()
	for i := 0; i < len(sizeDirPairs); i++ {
		if sizeDirPairs[i].Size+currentSpace > neededSpace {
			return strconv.Itoa(sizeDirPairs[i].Size)
		}
	}
	panic("Couldn't find any dir that is big enough.")
}

//
//    Now, you're ready to choose a directory to delete.
//
//    The total disk space available to the filesystem is 70000000. To run
//    the update, you need unused space of at least 30000000. You need to
//    find a directory you can delete that will free up enough space to run
//    the update.
//
//    In the example above, the total size of the outermost directory (and
//    thus the total amount of used space) is 48381165; this means that the
//    size of the unused space must currently be 21618835, which isn't quite
//    the 30000000 required by the update. Therefore, the update still
//    requires a directory with total size of at least 8381165 to be deleted
//    before it can run.
//
//    To achieve this, you have the following options:
//      * Delete directory e, which would increase unused space by 584.
//      * Delete directory a, which would increase unused space by 94853.
//      * Delete directory d, which would increase unused space by 24933642.
//      * Delete directory /, which would increase unused space by 48381165.
//
//    Directories e and a are both too small; deleting them would not free up
//    enough space. However, directories d and / are both big enough! Between
//    these, choose the smallest: d, increasing unused space by 24933642.
//
//    Find the smallest directory that, if deleted, would free up enough
//    space on the filesystem to run the update. What is the total size of
//    that directory?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
