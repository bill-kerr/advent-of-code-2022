package day05

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type Stack struct {
	crates []rune
	id int
}

func (s *Stack) Add(crates []rune) {
	s.crates = append(s.crates, crates...)
}

func (s *Stack) Remove(count int) []rune {
	if len(s.crates) - count < 0 {
		panic("Attempted to remove too many crates from the stack!")
	}

	removed := s.crates[len(s.crates) - count:]
	s.crates = s.crates[0:len(s.crates) - count]

	return removed
}

func (s *Stack) ToString() string {
	str := ""

	for i := len(s.crates) - 1; i >= 0; i-- {
		str += fmt.Sprintf("[%s]\n", string(s.crates[i]))
	}

	str += fmt.Sprintf(" %v ", s.id)

	return str
}

func (s *Stack) GetTopCrate() rune {
	return s.crates[len(s.crates) - 1]
}

func newStack(id int) *Stack {
	return &Stack{id: id, crates: []rune{}}
}

type Instruction struct {
	quantity int
	from int
	to int
}

func (i *Instruction) ToString() string {
	return fmt.Sprintf("Move %v from %v to %v", i.quantity, i.from, i.to)
}

func newInstruction(quantity int, from int, to int) *Instruction {
	return &Instruction{quantity: quantity, from: from, to: to}
}

func splitLines(lines []string) ([]string, []string) {
	stackLines := []string{}
	instructionLines := []string{}

	hasSeenSplit := false
	for _, line := range lines {
		if line == "" {
			hasSeenSplit = true
			continue
		}

		if !hasSeenSplit {
			stackLines = append(stackLines, line)
		} else {
			instructionLines = append(instructionLines, line)
		}
	}

	util.Reverse(stackLines)

	return stackLines, instructionLines
}

func parseStackLines(stackLines []string) map[int]*Stack {
	idLine := strings.ReplaceAll(stackLines[0], " ", "")
	stacks := map[int]*Stack{}

	for _, id := range idLine {
		stacks[util.Rtoi(id)] = newStack(util.Rtoi(id))
	}

	crateLines := stackLines[1:]
	for _, line := range crateLines {
		crateNames := []rune{}
		for i := 1; i < len(line); i += 4 {
				crateNames = append(crateNames, rune(line[i]))
		}

		for i, name := range crateNames {
			if name != ' ' {
				stacks[i + 1].Add([]rune{name})
			}
		}
	}

	return stacks
}

func parseInstructionLines(instructionLines []string) []*Instruction {
	instructions := []*Instruction{}

	for _, line := range instructionLines {
		withoutMove := strings.ReplaceAll(line, "move ", "")
		withoutFrom := strings.ReplaceAll(withoutMove, "from ", "")
		withoutTo := strings.ReplaceAll(withoutFrom, "to ", "")

		nums := strings.Split(withoutTo, " ")
		quantity := util.Atoi(nums[0])
		from := util.Atoi(nums[1])
		to := util.Atoi(nums[2])

		instructions = append(instructions, newInstruction(quantity, from, to))
	}

	return instructions
}

func part1(lines []string) {
	stackLines, instructionLines := splitLines(lines)

	stacks := parseStackLines(stackLines)

	instructions := parseInstructionLines(instructionLines)

	for _, instruction := range instructions {
		sourceStack := stacks[instruction.from]
		destinationStack := stacks[instruction.to]
		
		for i := 0; i < instruction.quantity; i++ {
			removed := sourceStack.Remove(1)
			destinationStack.Add(removed)
		}
	}

	sortedStacks := []*Stack{}
	for _, stack := range stacks {
		sortedStacks = append(sortedStacks, stack)
	}
	sort.Slice(sortedStacks, func (i, j int) bool {
		return sortedStacks[i].id < sortedStacks[j].id
	})

	var answer string
	for _, stack := range sortedStacks {
		answer += string(stack.GetTopCrate())
	}

	fmt.Println(answer)
}

func part2(lines []string) {
	stackLines, instructionLines := splitLines(lines)

	stacks := parseStackLines(stackLines)

	instructions := parseInstructionLines(instructionLines)

	for _, instruction := range instructions {
		sourceStack := stacks[instruction.from]
		destinationStack := stacks[instruction.to]
		
		removed := sourceStack.Remove(instruction.quantity)
		destinationStack.Add(removed)
	}

	sortedStacks := []*Stack{}
	for _, stack := range stacks {
		sortedStacks = append(sortedStacks, stack)
	}
	sort.Slice(sortedStacks, func (i, j int) bool {
		return sortedStacks[i].id < sortedStacks[j].id
	})

	var answer string
	for _, stack := range sortedStacks {
		answer += string(stack.GetTopCrate())
	}

	fmt.Println(answer)
}

func Run() {
	lines := util.OpenAndRead("./day05/input.txt")

	part1(lines)
	part2(lines)
}
