package day11

import (
	"fmt"
	"sort"
)

type Monkee struct {
	Items []int
	Operation func(old int) int
	Divisor int
	Targets [2]int
	Inspections int
}

func (m *Monkee) Catch(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkee) Throw(item int, target *Monkee) {
	target.Catch(item)
	m.Items = m.Items[1:]
}

func (m *Monkee) Inspect() int {
	m.Inspections++
	return m.Operation(m.Items[0])
}

func (m *Monkee) GetTarget(worryLevel int) int {
	if worryLevel % m.Divisor == 0 {
		return m.Targets[0]
	}

	return m.Targets[1]
}

func simulateRounds(monkees []*Monkee, numRounds int, afterInspect func(old int) int) int {
	round := 1

	for round <= numRounds {
		for _, monkee := range monkees {
			for len(monkee.Items) > 0 {
				worryLevel := monkee.Inspect()
	
				after := afterInspect(worryLevel)

				targetIndex := monkee.GetTarget(after)
				targetMonkee := monkees[targetIndex]
	
				monkee.Throw(after, targetMonkee)
			}
		}

		round++
	}

	sort.Slice(monkees, func (i, j int) bool {
		return monkees[i].Inspections > monkees[j].Inspections
	})

	return monkees[0].Inspections * monkees[1].Inspections
}

func part1(monkees []*Monkee) {
	monkeyBusiness := simulateRounds(monkees, 20, func(old int) int { return old / 3 })
	fmt.Println(monkeyBusiness)
}

func part2(monkees []*Monkee) {
	monkeyBusiness := simulateRounds(monkees, 10000, func(old int) int { return old })
	fmt.Println(monkeyBusiness)
}

func Run() {
	part1(Monkees)
	// part2(SampleMonkees)
}
