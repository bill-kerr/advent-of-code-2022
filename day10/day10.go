package day10

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type InstructionName string

const (
	AddX InstructionName = "addx"
	Noop InstructionName = "noop"
)

type Instruction struct {
	Name InstructionName
	Value int
	Duration int
}

func newInstruction(name InstructionName, value int) *Instruction {
	if name == Noop {
		return &Instruction{Name: Noop, Value: value, Duration: 1}
	}

	return &Instruction{Name: AddX, Value: value, Duration: 2}
}

func shouldCheckCycle(cycle int) bool {
	return (cycle + 20) % 40 == 0
}

func part1(instructions []*Instruction) {
	cycle := 1
	register := 1
	sumSignals := 0

	for _, instruction := range instructions {
		for j := 0; j < instruction.Duration; j++ {
			if shouldCheckCycle(cycle) {
				signalStrength := cycle * register
				sumSignals += signalStrength
			}

			cycle++
		}
		if instruction.Name == AddX {
			register += instruction.Value
		}
	}

	fmt.Println(sumSignals)
}

func part2(instructions []*Instruction) {
	register := 1
	crt := 0
	display := ""

	for _, instruction := range instructions {
		for j := 0; j < instruction.Duration; j++ {
			if crt % 40 <= register + 1 && crt % 40 >= register - 1 {
				display += "#"
			} else {
				display += "."
			}
			
			crt++
			if crt % 40 == 0 {
				display += "\n"
			}
		}
		if instruction.Name == AddX {
			register += instruction.Value
		}
	}

	fmt.Println(display)
}

func parseInput(lines []string) []*Instruction {
	instructions := make([]*Instruction, len(lines))

	for i, line := range lines {
		name := line[0:4]

		if name == "addx" {
			instructions[i] = newInstruction(AddX, util.Atoi(line[5:]))
		}

		if name == "noop" {
			instructions[i] = newInstruction(Noop, 0)
		}
	}

	return instructions
}

func Run() {
	lines := util.OpenAndRead("./day10/input.txt")

	instructions := parseInput(lines)

	part1(instructions)
	part2(instructions)
}
