package day09

import (
	"fmt"
	"strings"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	X int
	Y int
}

func newPosition(x, y int) *Position {
	return &Position{X: x, Y: y}
}

func (p *Position) ToString() string {
	return fmt.Sprintf("%v, %v", p.X, p.Y)
}

func (p *Position) IsTouching(otherPosition *Position) bool {
	inX := p.X >= otherPosition.X - 1 && p.X <= otherPosition.X + 1
	inY := p.Y >= otherPosition.Y - 1 && p.Y <= otherPosition.Y + 1

	return inX && inY
}

func (p *Position) DirectionTo(otherPosition *Position) *Position {
	x := 0
	y := 0

	if p.X > otherPosition.X {
		x = -1
	} else if p.X < otherPosition.X {
		x = 1
	}

	if p.Y > otherPosition.Y {
		y = -1
	} else if p.Y < otherPosition.Y {
		y = 1
	}

	return newPosition(x, y)
}

type Movement struct {
	Direction Direction
	Distance int
}

type Rope struct {
	Knots []*Position
	TailPositions *util.Set[string]
}

func newRope(numKnots int) *Rope {
	knots := make([]*Position, numKnots)

	for i := 0; i < numKnots; i++ {
		knots[i] = newPosition(0, 0)
	}

	tailPositions := util.NewSet[string]()
	tailPositions.Add(knots[len(knots) - 1].ToString())

	return &Rope{Knots: knots, TailPositions: tailPositions}
}

func (r *Rope) Move(movement *Movement) {
	unitVector := 1
	if movement.Direction == Down || movement.Direction == Left {
		unitVector = -1
	}

	for i := 0; i < movement.Distance; i++ {
		// Move the head
		if movement.Direction == Right || movement.Direction == Left {
			r.Knots[0].X += unitVector
		} else {
			r.Knots[0].Y += unitVector
		}

		// For each knot that is not the head, make sure it's touching the one before it
		for j := 1; j < len(r.Knots); j++ {
			for !r.Knots[j].IsTouching(r.Knots[j - 1]) {
				tailMovement := r.Knots[j].DirectionTo(r.Knots[j - 1])
				r.Knots[j].X += tailMovement.X
				r.Knots[j].Y += tailMovement.Y
	
				// If this is the tail of the rope, update its visited positions
				if j == len(r.Knots) - 1 {
					r.TailPositions.Add(r.Knots[j].ToString())
				}
			}
		}
	}
}

func getDirection(str string) Direction {
	switch str {
	case "R":
		return Right
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	}

	panic("Invalid direction!")
}

func parseMovements(lines []string) []*Movement {
	movements := make([]*Movement, len(lines))

	for i, line := range lines {
		info := strings.Split(line, " ")

		movements[i] = &Movement{Direction: getDirection(info[0]), Distance: util.Atoi(info[1])}
	}

	return movements
}

func part1(movements []*Movement) {
	rope := newRope(2)

	for _, movement := range movements {
		rope.Move(movement)
	}

	fmt.Println(rope.TailPositions.Size())
}

func part2(movements []*Movement) {
	rope := newRope(10)

	for _, movement := range movements {
		rope.Move(movement)
	}

	fmt.Println(rope.TailPositions.Size())
}

func Run() {
	lines := util.OpenAndRead("./day09/input.txt")

	movements := parseMovements(lines)

	part1(movements)
	part2(movements)
}
