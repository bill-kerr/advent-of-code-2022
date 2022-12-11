package day11

import (
	"fmt"
	"sort"
)

type Operator int

const (
	Add Operator = iota
	Multiply
	Square
)

type Operation struct {
	Operator Operator
	Magnitude int
}

type Monkey struct {
	Items []int
	Operation *Operation
	Divisor int
	Targets [2]int
	Inspections int
}

func (m *Monkey) Catch(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) Throw(target *Monkey) {
	target.Catch(m.Items[0])
	m.Items = m.Items[1:]
}

func (m *Monkey) Inspect() int {
	m.Inspections++

	if m.Operation.Operator == Add {
		return m.Items[0] + m.Operation.Magnitude
	} else if m.Operation.Operator == Multiply {
		return m.Items[0] * m.Operation.Magnitude
	}
	return m.Items[0] * m.Items[0]
}

func (m *Monkey) GetTarget(afterInspect func(worryLevel int) int) int {
	m.Inspections++

	worryLevel := 0
	if m.Operation.Operator == Add {
		worryLevel = m.Items[0] + m.Operation.Magnitude
	} else if m.Operation.Operator == Multiply {
		worryLevel = m.Items[0] * m.Operation.Magnitude
	} else {
		worryLevel = m.Items[0] * m.Items[0]
	}

	newWorryLevel := afterInspect(worryLevel)

	m.Items[0] = newWorryLevel

	if newWorryLevel % m.Divisor == 0 {
		return m.Targets[0]
	}

	return m.Targets[1]
}

func simulateRounds(monkeys []*Monkey, numRounds int, afterInspect func(old int) int) int {
	round := 1

	for round <= numRounds {
		for _, monkey := range monkeys {
			for len(monkey.Items) > 0 {
				targetIndex := monkey.GetTarget(afterInspect)
				targetMonkey := monkeys[targetIndex]
				monkey.Throw(targetMonkey)
			}
		}

		round++
	}

	sort.Slice(monkeys, func (i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	return monkeys[0].Inspections * monkeys[1].Inspections
}

func part1(monkees []*Monkey) {
	monkeyBusiness := simulateRounds(monkees, 20, func(old int) int { return old / 3 })
	fmt.Println(monkeyBusiness)
}

func part2(monkees []*Monkey) {
	monkeyBusiness := simulateRounds(monkees, 10000, func(old int) int { return old })
	fmt.Println(monkeyBusiness)
}

func Run() {
	part1(Monkeys)
	// part2(SampleMonkees)
}
