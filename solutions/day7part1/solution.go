package day7part1

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

type File struct {
	name string
	size int
}

type Directory struct {
	name           string
	parent         *Directory
	subdirectories []*Directory
	files          []*File
}

func newDirectory(name string) *Directory {
	return &Directory{
		name:           name,
		subdirectories: []*Directory{},
		files:          []*File{},
	}
}

func (d *Directory) addSubdirectory(name string) {
	newDir := newDirectory(name)
	newDir.parent = d
	d.subdirectories = append(d.subdirectories, newDir)
}

func (d *Directory) getSubdirectory(name string) *Directory {
	for _, dir := range d.subdirectories {
		if dir.name == name {
			return dir
		}
	}
	panic("Subdirectory not found.")
}

func (d *Directory) GetTotalSize() int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, dir := range d.subdirectories {
		size += dir.GetTotalSize()
	}
	return size
}

func (d *Directory) addFile(file *File) {
	d.files = append(d.files, file)
}

type SizeDirPair struct {
	Size int
	dir  *Directory
}
type SizeDirPairList []SizeDirPair

func (s SizeDirPairList) Less(i, j int) bool {
	return s[i].Size < s[j].Size
}

func (s SizeDirPairList) Len() int {
	return len(s)
}

func (s SizeDirPairList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func parseFile(line []string) *File {
	size, err := strconv.Atoi(line[0])
	if err != nil {
		log.Fatal(err)
	}
	return &File{
		size: size,
		name: line[1],
	}
}

func ParseDirectoryTree(input string) *Directory {
	lines := strings.Split(input, "\n")
	currentDirectory := newDirectory("/")
	root := currentDirectory
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " ")
		switch split[0] {
		case "$":
			if split[1] == "cd" {
				if split[2] == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					parent := currentDirectory
					currentDirectory = parent.getSubdirectory(split[2])
				}
			}
		case "dir":
			currentDirectory.addSubdirectory(split[1])
		default:
			currentDirectory.addFile(parseFile(split))
		}
	}
	return root
}

func GetTotalSizes(tree *Directory) []SizeDirPair {
	pairs := []SizeDirPair{
		{
			Size: tree.GetTotalSize(),
			dir:  tree,
		},
	}
	for _, sub := range tree.subdirectories {
		pairs = append(pairs, GetTotalSizes(sub)...)
	}
	return pairs
}

func (Solver) Solve(input string) string {
	tree := ParseDirectoryTree(input)
	sizeDirPairs := GetTotalSizes(tree)
	sort.Sort(SizeDirPairList(sizeDirPairs))
	sum := 0
	for i := 0; i < len(sizeDirPairs); i++ {
		if sizeDirPairs[i].Size > 100000 {
			break
		}
		sum += sizeDirPairs[i].Size
	}
	return strconv.Itoa(sum)
}

//
//    You can hear birds chirping and raindrops hitting leaves as the
//    expedition proceeds. Occasionally, you can even hear much louder sounds
//    in the distance; how big do the animals get out here, anyway?
//
//    The device the Elves gave you has problems with more than just its
//    communication system. You try to run a system update:
// $ system-update --please --pretty-please-with-sugar-on-top
// Error: No space left on device
//
//    Perhaps you can delete some files to make space for the update?
//
//    You browse around the filesystem to assess the situation and save the
//    resulting terminal output (your puzzle input). For example:
// $ cd /
// $ ls
// dir a
// 14848514 b.txt
// 8504156 c.dat
// dir d
// $ cd a
// $ ls
// dir e
// 29116 f
// 2557 g
// 62596 h.lst
// $ cd e
// $ ls
// 584 i
// $ cd ..
// $ cd ..
// $ cd d
// $ ls
// 4060174 j
// 8033020 d.log
// 5626152 d.ext
// 7214296 k
//
//    The filesystem consists of a tree of files (plain data) and directories
//    (which can contain other directories or files). The outermost directory
//    is called /. You can navigate around the filesystem, moving into or out
//    of directories and listing the contents of the directory you're
//    currently in.
//
//    Within the terminal output, lines that begin with $ are commands you
//    executed, very much like some modern computers:
//      * cd means change directory. This changes which directory is the
//        current directory, but the specific result depends on the argument:
//           + cd x moves in one level: it looks in the current directory for
//             the directory named x and makes it the current directory.
//           + cd .. moves out one level: it finds the directory that
//             contains the current directory, then makes that directory the
//             current directory.
//           + cd / switches the current directory to the outermost
//             directory, /.
//      * ls means list. It prints out all of the files and directories
//        immediately contained by the current directory:
//           + 123 abc means that the current directory contains a file named
//             abc with size 123.
//           + dir xyz means that the current directory contains a directory
//             named xyz.
//
//    Given the commands and output in the example above, you can determine
//    that the filesystem looks visually like this:
// - / (dir)
//   - a (dir)
//     - e (dir)
//       - i (file, size=584)
//     - f (file, size=29116)
//     - g (file, size=2557)
//     - h.lst (file, size=62596)
//   - b.txt (file, size=14848514)
//   - c.dat (file, size=8504156)
//   - d (dir)
//     - j (file, size=4060174)
//     - d.log (file, size=8033020)
//     - d.ext (file, size=5626152)
//     - k (file, size=7214296)
//
//    Here, there are four directories: / (the outermost directory), a and d
//    (which are in /), and e (which is in a). These directories also contain
//    files of various sizes.
//
//    Since the disk is full, your first step should probably be to find
//    directories that are good candidates for deletion. To do this, you need
//    to determine the total size of each directory. The total size of a
//    directory is the sum of the sizes of the files it contains, directly or
//    indirectly. (Directories themselves do not count as having any
//    intrinsic size.)
//
//    The total sizes of the directories above can be found as follows:
//      * The total size of directory e is 584 because it contains a single
//        file i of size 584 and no other directories.
//      * The directory a has total size 94853 because it contains files f
//        (size 29116), g (size 2557), and h.lst (size 62596), plus file i
//        indirectly (a contains e which contains i).
//      * Directory d has total size 24933642.
//      * As the outermost directory, / contains every file. Its total size
//        is 48381165, the sum of the size of every file.
//
//    To begin, find all of the directories with a total size of at most
//    100000, then calculate the sum of their total sizes. In the example
//    above, these directories are a and e; the sum of their total sizes is
//    95437 (94853 + 584). (As in this example, this process can count files
//    more than once!)
//
//    Find all of the directories with a total size of at most 100000. What
//    is the sum of the total sizes of those directories?
//
//    To begin, get your puzzle input.
//
//    Answer: ____________________ [Submit]
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
