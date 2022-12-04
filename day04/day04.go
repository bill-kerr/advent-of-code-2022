package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type Assignment struct {
	start int
	end   int
}

func newAssignment(start int, end int) *Assignment {
	a := &Assignment{start: start, end: end}
	return a
}

func (a *Assignment) FullyContains(other *Assignment) bool {
	return a.start <= other.start && a.end >= other.end
}

func (a *Assignment) InRange(val int) bool {
	return a.start <= val && a.end >= val
}

type AssignmentPair struct {
	a *Assignment
	b *Assignment
}

func newAssignmentPair(a *Assignment, b *Assignment) *AssignmentPair {
	p := &AssignmentPair{a: a, b: b}
	return p
}

func (p *AssignmentPair) HasFullyContainedAssignment() bool {
	return p.a.FullyContains(p.b) || p.b.FullyContains(p.a)
}

func (p *AssignmentPair) HasOverlap() bool {
	return p.a.InRange(p.b.start) || p.a.InRange(p.b.end) || p.b.InRange(p.a.start) || p.b.InRange(p.a.end)
}

func makeAssignment(assignment string) *Assignment {
	assignmentRange := strings.Split(assignment, "-")

	start, err := strconv.Atoi(assignmentRange[0])
	if err != nil {
		log.Fatal("Failed to parse assignment range")
	}

	end, err := strconv.Atoi(assignmentRange[1])
	if err != nil {
		log.Fatal("Failed to parse assignment range")
	}

	return newAssignment(start, end)
}

func makeAssignmentPair(line string) *AssignmentPair {
	pairs := strings.Split(line, ",")

	a := makeAssignment(pairs[0])
	b := makeAssignment(pairs[1])

	return newAssignmentPair(a, b)
}

func part1(lines []string) {
	count := 0
	for _, line := range lines {
		pair := makeAssignmentPair(line)
		if pair.HasFullyContainedAssignment() {
			count++
		}
	}

	fmt.Println(count)
}

func part2(lines []string) {
	count := 0
	for _, line := range lines {
		pair := makeAssignmentPair(line)
		if pair.HasOverlap() {
			count++
		}
	}

	fmt.Println(count)
}

func Run() {
	lines := util.OpenAndRead("./day04/input.txt")

	part1(lines)
	part2(lines)
}