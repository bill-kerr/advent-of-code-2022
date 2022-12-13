package day11

func newItemValue(value int, size int) int {
	itemValue := make([]int, size)
	for i := 0; i < size; i++ {
		itemValue[i] = value
	}
	return value
}

var monkey0 *Monkey = &Monkey{
	Items: []int{
		newItemValue(97, 8),
		newItemValue(81, 8),
		newItemValue(57, 8),
		newItemValue(57, 8),
		newItemValue(91, 8),
		newItemValue(61, 8),
	},
	Operation: &Operation{Operator: Multiply, Magnitude: 7},
	Divisor:   11,
	Targets:   [2]int{5, 6},
}

var monkey1 *Monkey = &Monkey{
	Items: []int{
		newItemValue(88, 8),
		newItemValue(62, 8),
		newItemValue(68, 8),
		newItemValue(90, 8),
	},
	Operation: &Operation{Operator: Multiply, Magnitude: 17},
	Divisor:   19,
	Targets:   [2]int{4, 2},
}

var monkey2 *Monkey = &Monkey{
	Items: []int{
		newItemValue(74, 8),
		newItemValue(87, 8),
	},
	Operation: &Operation{Operator: Add, Magnitude: 2},
	Divisor:   5,
	Targets:   [2]int{7, 4},
}

var monkey3 *Monkey = &Monkey{
	Items: []int{
		newItemValue(53, 8),
		newItemValue(81, 8),
		newItemValue(60, 8),
		newItemValue(87, 8),
		newItemValue(90, 8),
		newItemValue(99, 8),
		newItemValue(75, 8),
	},
	Operation: &Operation{Operator: Add, Magnitude: 1},
	Divisor:   2,
	Targets:   [2]int{2, 1},
}

var monkey4 *Monkey = &Monkey{
	Items:     []int{newItemValue(57, 8)},
	Operation: &Operation{Operator: Add, Magnitude: 6},
	Divisor:   13,
	Targets:   [2]int{7, 0},
}

var monkey5 *Monkey = &Monkey{
	Items: []int{
		newItemValue(54, 8),
		newItemValue(84, 8),
		newItemValue(91, 8),
		newItemValue(55, 8),
		newItemValue(59, 8),
		newItemValue(72, 8),
		newItemValue(75, 8),
		newItemValue(70, 8),
	},
	Operation: &Operation{Operator: Square, Magnitude: 0},
	Divisor:   7,
	Targets:   [2]int{6, 3},
}

var monkey6 *Monkey = &Monkey{
	Items: []int{
		newItemValue(95, 8),
		newItemValue(79, 8),
		newItemValue(79, 8),
		newItemValue(68, 8),
		newItemValue(78, 8),
	},
	Operation: &Operation{Operator: Add, Magnitude: 3},
	Divisor:   3,
	Targets:   [2]int{1, 3},
}

var monkey7 *Monkey = &Monkey{
	Items: []int{
		newItemValue(61, 8),
		newItemValue(97, 8),
		newItemValue(67, 8),
	},
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
	Items:     []int{newItemValue(79, 4), newItemValue(98, 4)},
	Operation: &Operation{Operator: Multiply, Magnitude: 19},
	Divisor:   23,
	Targets:   [2]int{2, 3},
}

var sampleMonkey1 *Monkey = &Monkey{
	Items: []int{
		newItemValue(54, 4),
		newItemValue(65, 4),
		newItemValue(75, 4),
		newItemValue(74, 4),
	},
	Operation: &Operation{Operator: Add, Magnitude: 6},
	Divisor:   19,
	Targets:   [2]int{2, 0},
}

var sampleMonkey2 *Monkey = &Monkey{
	Items: []int{
		newItemValue(79, 4),
		newItemValue(60, 4),
		newItemValue(97, 4),
	},
	Operation: &Operation{Operator: Square, Magnitude: 0},
	Divisor:   13,
	Targets:   [2]int{1, 3},
}

var sampleMonkey3 *Monkey = &Monkey{
	Items:     []int{newItemValue(74, 4)},
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
