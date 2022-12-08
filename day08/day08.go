package day08

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type Tree struct {
	Height int
	X int
	Y int
	Visible bool
	ScenicScore int
}

func newTree(height, x, y int, row, col []int) *Tree {
	tree := &Tree{Height: height, X: x, Y: y}
	blockedDirections := 0

	scoreTop := 0
	scoreBottom := 0
	scoreLeft := 0
	scoreRight := 0

	// If it's on the edge, we know it's visible and has a scenic score of 0
	if x == 0 || y == 0 || x == len(row) - 1 || y == len(col) - 1 {
		tree.Visible = true
		tree.ScenicScore = 0
		return tree
	}

	// Visible from left?
	for i := x - 1; i >= 0; i-- {
		scoreLeft++

		if row[i] >= height {
			blockedDirections++
			break
		}
	}

	// Visible from right?
	for i := x + 1; i < len(row); i++ {
		scoreRight++

		if row[i] >= height {
			blockedDirections++
			break
		}
	}

	// Visible from top?
	for i := y - 1; i >= 0; i-- {
		scoreTop++

		if col[i] >= height {
			blockedDirections++
			break
		}
	}

	// Visible from bottom?
	for i := y + 1; i < len(col); i++ {
		scoreBottom++

		if col[i] >= height {
			blockedDirections++
			break
		}
	}

	tree.Visible = blockedDirections < 4
	tree.ScenicScore = scoreBottom * scoreTop * scoreRight * scoreLeft
	return tree
}

func getCol(ints [][]int, col int) []int {
	column := make([]int, len(ints))

	for i := 0; i < len(ints); i++ {
		column[i] = ints[i][col]
	}

	return column
}

func part1(input [][]int) {
	visibleCount := 0

	for y, row := range input {
		for x, height := range row {
			col := getCol(input, x)
			tree := newTree(height, x, y, row, col)

			if tree.Visible {
				visibleCount++
			}
		}
	}

	fmt.Println(visibleCount)
}

func part2(input [][]int) {
	highestScore := 0

	for y, row := range input {
		for x, height := range row {
			col := getCol(input, x)
			tree := newTree(height, x, y, row, col)

			if tree.ScenicScore > highestScore {
				highestScore = tree.ScenicScore
			}
		}
	}

	fmt.Println(highestScore)
}

func transformLines(lines []string) [][]int {
	ints := make([][]int, len(lines))
	
	for y, line := range lines {
		for _, char := range line {
			height := util.Atoi(string(char))
			ints[y] = append(ints[y], height)
		}
	}

	return ints
}

func Run() {
	lines := util.OpenAndRead("./day08/input.txt")
	ints := transformLines(lines)

	part1(ints)
	part2(ints)
}
