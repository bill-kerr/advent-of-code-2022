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

func (m *Monkey) Op() {
	if m.Operation.Operator == Add {
		m.Items[0] += m.Operation.Magnitude
	} else if m.Operation.Operator == Multiply {
		m.Items[0] *= m.Operation.Magnitude
	} else {
		m.Items[0] *= m.Items[0]
	}
}

func (m *Monkey) Throw(target *Monkey) {
	item := m.Items[0]
	m.Items = m.Items[1:]
	target.Catch(item)
}

func (m *Monkey) Catch(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) GetTargetIndex() int {
	if m.Items[0] % m.Divisor == 0 {
		return m.Targets[0]
	}
	return m.Targets[1]
}

func simulateRounds(monkeys []*Monkey, numRounds int, onInspect func (val int, product int) int) int {
	round := 1

	product := 1
	for _, monkey := range monkeys {
		product *= monkey.Divisor
	}

	for round <= numRounds {
		for _, monkey := range monkeys {
			for len(monkey.Items) > 0 {
				monkey.Inspections++
				
				monkey.Op()
				monkey.Items[0] = onInspect(monkey.Items[0], product)

				targetMonkey := monkeys[monkey.GetTargetIndex()]
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
	monkeyBusiness := simulateRounds(monkees, 20, func(val int, product int) int {
		return val / 3
	})
	fmt.Println(monkeyBusiness)
}

func part2(monkees []*Monkey) {
	monkeyBusiness := simulateRounds(monkees, 10000, func(val int, product int) int { 
		return val % product 
	})
	fmt.Println(monkeyBusiness)
}

func Run() {
	part1(Monkeys)
	part2(Monkeys)
}
