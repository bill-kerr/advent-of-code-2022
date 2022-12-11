package day11

var monkey0 *Monkey = &Monkey{
	Items: []int{
		97,
		81,
		57,
		57,
		91,
		61,
	},
	Operation: &Operation{Operator: Multiply, Magnitude: 7},
	Divisor:   11,
	Targets:   [2]int{5, 6},
}

var monkey1 *Monkey = &Monkey{
	Items: []int{
		88,
		62,
		68,
		90,
	},
	Operation: &Operation{Operator: Multiply, Magnitude: 17},
	Divisor:   19,
	Targets:   [2]int{4, 2},
}

var monkey2 *Monkey = &Monkey{
	Items:     []int{74, 87},
	Operation: &Operation{Operator: Add, Magnitude: 2},
	Divisor:   5,
	Targets:   [2]int{7, 4},
}

var monkey3 *Monkey = &Monkey{
	Items: []int{
		53,
		81,
		60,
		87,
		90,
		99,
		75,
	},
	Operation: &Operation{Operator: Add, Magnitude: 1},
	Divisor:   2,
	Targets:   [2]int{2, 1},
}

var monkey4 *Monkey = &Monkey{
	Items:     []int{57},
	Operation: &Operation{Operator: Add, Magnitude: 6},
	Divisor:   13,
	Targets:   [2]int{7, 0},
}

var monkey5 *Monkey = &Monkey{
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
	Operation: &Operation{Operator: Square, Magnitude: 0},
	Divisor:   7,
	Targets:   [2]int{6, 3},
}

var monkey6 *Monkey = &Monkey{
	Items: []int{
		95,
		79,
		79,
		68,
		78,
	},
	Operation: &Operation{Operator: Add, Magnitude: 3},
	Divisor:   3,
	Targets:   [2]int{1, 3},
}

var monkey7 *Monkey = &Monkey{
	Items:     []int{61, 97, 67},
	Operation: &Operation{Operator: Add, Magnitude: 4},
	Divisor:   17,
	Targets:   [2]int{0, 5},
}

var Monkeys []*Monkey = []*Monkey{
	monkey0,
	monkey1,
	monkey2,
	monkey3,
	monkey4,
	monkey5,
	monkey6,
	monkey7,
}

var sampleMonkey0 *Monkey = &Monkey{
	Items:     []int{79, 98},
	Operation: &Operation{Operator: Multiply, Magnitude: 19},
	Divisor:   23,
	Targets:   [2]int{2, 3},
}

var sampleMonkey1 *Monkey = &Monkey{
	Items:     []int{54, 65, 75, 74},
	Operation: &Operation{Operator: Add, Magnitude: 6},
	Divisor:   19,
	Targets:   [2]int{2, 0},
}

var sampleMonkey2 *Monkey = &Monkey{
	Items:     []int{79, 60, 97},
	Operation: &Operation{Operator: Square, Magnitude: 0},
	Divisor:   13,
	Targets:   [2]int{1, 3},
}

var sampleMonkey3 *Monkey = &Monkey{
	Items:     []int{74},
	Operation: &Operation{Operator: Add, Magnitude: 3},
	Divisor:   17,
	Targets:   [2]int{0, 1},
}

var SampleMonkeys []*Monkey = []*Monkey{
	sampleMonkey0,
	sampleMonkey1,
	sampleMonkey2,
	sampleMonkey3,
}