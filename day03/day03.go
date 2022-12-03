package day03

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

func splitString(str string) ([]rune, []rune) {
	half := int(len(str) / 2)
	return []rune(str[0:half]), []rune(str[half:])
}

func getPriorityValue(r rune) int {
	i := int(r)
	if i >= 97 {
		return i - 96
	}
	return i - 38
}

func getPriorities(items []rune) []int {
	priorities := []int{}

	for _, val := range items {
		priorities = append(priorities, getPriorityValue(val))
	}

	return priorities
}

func Run() {
	part1()
	part2()
}

func part1() {
	lines := util.OpenAndRead("./day03/input.txt")

	prioritySum := 0

	for _, val := range lines {
		first, second := splitString(val)
		
		firstPriorities := getPriorities(first)
		secondPriorities := getPriorities(second)

		matching := 0
		for _, f := range firstPriorities {
			for _, s := range secondPriorities {
				if f == s {
					matching = f
				}
			}
		}

		prioritySum += matching
	}

	fmt.Println(prioritySum)
}

func getPrioritySet(contents string) *util.Set[int] {
	set := util.NewSet[int]()

	for _, c := range contents {
		set.Add(getPriorityValue(c))
	}

	return set
}

func part2() {
	lines := util.OpenAndRead("./day03/input.txt")

	prioritySum := 0

	group := []string{}

	for _, val := range lines {
		if len(group) == 3 {
			group = []string{}
		}
		group = append(group, val)

		if (len(group) == 3) {
			set1 := getPrioritySet(group[0])
			set2 := getPrioritySet(group[1])
			set3 := getPrioritySet(group[2])

			for priority := range set1.ToMap() {
				if set2.Has(priority) && set3.Has(priority) {
					prioritySum += priority
				}
			}
		}
	}

	fmt.Println(prioritySum)
}