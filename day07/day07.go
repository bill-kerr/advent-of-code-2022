package day07

import (
	"fmt"
	"strings"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type Command int

const (
	CD Command = iota
	LS
)

type File struct {
	name string
	size int
}

func newFile(name string, size int) *File {
	return &File{name: name, size: size}
}

type Directory struct {
	parent *Directory
	files []*File
	dirs []*Directory
	name string
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name, parent: parent}
}

func (d *Directory) AddDirectory(name string) *Directory {
	// Retrun early if that directory already exists
	for _, directory := range d.dirs {
		if directory.name == name {
			return directory
		}
	}

	newDir := newDirectory(name, d)
	d.dirs = append(d.dirs, newDir)
	return newDir
}

func (d *Directory) AddFile(name string, size int) *File {
	for _, file := range d.files {
		if file.name == name {
			return file
		}
	}

	f := newFile(name, size)
	d.files = append(d.files, f)
	return f
}

func (d *Directory) GetSize() int {
	dirSize := 0
	fileSize := 0

	for _, f := range d.files {
		fileSize += f.size
	}

	for _, dir := range d.dirs {
		dirSize += dir.GetSize()
	}

	return dirSize + fileSize
}

func (d *Directory) Traverse(onVisit func(directory *Directory)) {
	onVisit(d)

	for _, dir := range d.dirs {
		dir.Traverse(onVisit)
	}
}

func sumDirSizes(root *Directory, sizeLimit int) int {
	total := 0

	for _, dir := range root.dirs {
		dirSize := dir.GetSize()
		fmt.Println("dirSize for dir", dir.name, "is", dirSize)

		if dirSize <= sizeLimit {
			total += dirSize
		} 
			
		total += sumDirSizes(dir, sizeLimit)
	}

	return total
}

func parseInput(lines []string) *Directory {
	root := newDirectory("/", nil)
	current := root

	for _, line := range lines {
		if line[0] == '$' && line[2:4] == "cd" {
			dest := line[5:]

			// Check if destination is root directory. If not, add a new one
			if dest == root.name {
				current = root
			
			// Check if we're going backwards...if so, assign current to parent
			} else if dest == ".." {
				current = current.parent

			// Else, we need to add the directory since it may or may not exist yet
			} else {
				d := current.AddDirectory(dest)
				current = d
			}
		}

		// If it's not a command, we're listing files/dirs
		if line[0] != '$' {
			fileInfo := strings.Split(line, " ")

			if len(fileInfo[0]) == 3 && fileInfo[0] == "dir" {
				current.AddDirectory(fileInfo[1])
			} else {
				size := util.Atoi(fileInfo[0])
				name := fileInfo[1]
				current.AddFile(name, size)
			}
		}
	}

	return root
}

func part1(lines []string) {
	root := parseInput(lines)

	sum := sumDirSizes(root, 100000)
	fmt.Println(sum)
}

func part2(lines []string) {
	root := parseInput(lines)

	rootSize := root.GetSize()
	unusedSpace := 70000000 - rootSize
	requiredSpace := 30000000 - unusedSpace
	smallest := rootSize

	root.Traverse(func (d *Directory) {
		size := d.GetSize()

		fmt.Println(d.name, size)

		if (size <= smallest && size >= requiredSpace) {
			smallest = size
		}
	})

	fmt.Println(smallest)
}

func Run() {
	lines := util.OpenAndRead("./day07/input.txt")

	part1(lines)
	part2(lines)
}
