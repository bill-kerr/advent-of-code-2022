package day11

var monkee0 *Monkee = &Monkee{
	Items: []int{
		97,
		81,
		57,
		57,
		91,
		61,
	},
	Operation: func(old int) int { return old * 7 },
	Divisor:   11,
	Targets:   [2]int{5, 6},
}

var monkee1 *Monkee = &Monkee{
	Items: []int{
		88,
		62,
		68,
		90,
	},
	Operation: func(old int) int { return old * 17 },
	Divisor:   19,
	Targets:   [2]int{4, 2},
}

var monkee2 *Monkee = &Monkee{
	Items:     []int{74, 87},
	Operation: func(old int) int { return old + 2 },
	Divisor:   5,
	Targets:   [2]int{7, 4},
}

var monkee3 *Monkee = &Monkee{
	Items: []int{
		53,
		81,
		60,
		87,
		90,
		99,
		75,
	},
	Operation: func(old int) int { return old + 1 },
	Divisor:   2,
	Targets:   [2]int{2, 1},
}

var monkee4 *Monkee = &Monkee{
	Items:     []int{57},
	Operation: func(old int) int { return old + 6 },
	Divisor:   13,
	Targets:   [2]int{7, 0},
}

var monkee5 *Monkee = &Monkee{
	Items: []int{
		54,
		84,
		91,
		55,
		59,
		72,
		75,
		70,
	},
	Operation: func(old int) int { return old * old },
	Divisor:   7,
	Targets:   [2]int{6, 3},
}

var monkee6 *Monkee = &Monkee{
	Items: []int{
		95,
		79,
		79,
		68,
		78,
	},
	Operation: func(old int) int { return old + 3 },
	Divisor:   3,
	Targets:   [2]int{1, 3},
}

var monkee7 *Monkee = &Monkee{
	Items:     []int{61, 97, 67},
	Operation: func(old int) int { return old + 4 },
	Divisor:   17,
	Targets:   [2]int{0, 5},
}

var Monkees []*Monkee = []*Monkee{
	monkee0,
	monkee1,
	monkee2,
	monkee3,
	monkee4,
	monkee5,
	monkee6,
	monkee7,
}

var sampleMonkee0 *Monkee = &Monkee{
	Items:     []int{79, 98},
	Operation: func(old int) int { return old * 19 },
	Divisor:   23,
	Targets:   [2]int{2, 3},
}

var sampleMonkee1 *Monkee = &Monkee{
	Items:     []int{54, 65, 75, 74},
	Operation: func(old int) int { return old + 6 },
	Divisor:   19,
	Targets:   [2]int{2, 0},
}

var sampleMonkee2 *Monkee = &Monkee{
	Items:     []int{79, 60, 97},
	Operation: func(old int) int { return old * old },
	Divisor:   13,
	Targets:   [2]int{1, 3},
}

var sampleMonkee3 *Monkee = &Monkee{
	Items:     []int{74},
	Operation: func(old int) int { return old + 3 },
	Divisor:   17,
	Targets:   [2]int{0, 1},
}

var SampleMonkees []*Monkee = []*Monkee{
	sampleMonkee0,
	sampleMonkee1,
	sampleMonkee2,
	sampleMonkee3,
}