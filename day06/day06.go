package day06

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

func shift(window []rune, newItem rune) []rune {
	shifted := make([]rune, len(window))
	copy(shifted, window[1:])
	shifted[len(window) - 1] = newItem
	return shifted
}

func cardinality(window []rune) int {
	set := util.NewSet[rune]()
	for _, v := range window {
		set.Add(v)
	}
	return set.Size()
}

func distinctOffset(stream string, uniqueLength int) int {
	window := make([]rune, uniqueLength)
	for i, char := range stream[:uniqueLength] {
		window[i] = char
	}

	offset := 0

	for i := uniqueLength; i < len(stream); i++ {
		if len(window) == uniqueLength {
			c := cardinality(window)

			if c == uniqueLength {
				offset = i
				break
			}
		}

		window = shift(window, rune(stream[i]))
	}

	return offset
}

func part1(stream string) {
	fmt.Println(distinctOffset(stream, 4))
}

func part2(stream string) {
	fmt.Println(distinctOffset(stream, 14))
}

func Run() {
	lines := util.OpenAndRead("./day06/input.txt")

	part1(lines[0])
	part2(lines[0])
}
