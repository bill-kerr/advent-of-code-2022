package day09

import (
	"fmt"
	"strings"

	"github.com/bill-kerr/advent-of-code-2022/util"
)


type Vector struct {
	X int
	Y int
}

func newVector(x, y int) *Vector {
	return &Vector{X: x, Y: y}
}

func (v *Vector) ToString() string {
	return fmt.Sprintf("%v, %v", v.X, v.Y)
}

func (v *Vector) IsTouching(otherPosition *Vector) bool {
	inX := v.X >= otherPosition.X - 1 && v.X <= otherPosition.X + 1
	inY := v.Y >= otherPosition.Y - 1 && v.Y <= otherPosition.Y + 1

	return inX && inY
}

func (v *Vector) IsEqual(otherVector *Vector) bool {
	return v.X == otherVector.X && v.Y == otherVector.Y
}

func (v *Vector) DirectionTo(otherPosition *Vector) *Vector {
	x := 0
	y := 0

	if v.X > otherPosition.X {
		x = -1
	} else if v.X < otherPosition.X {
		x = 1
	}

	if v.Y > otherPosition.Y {
		y = -1
	} else if v.Y < otherPosition.Y {
		y = 1
	}

	return newVector(x, y)
}

type Rope struct {
	Knots []*Vector
	TailPositions *util.Set[string]
}

func newRope(numKnots int) *Rope {
	knots := make([]*Vector, numKnots)

	for i := 0; i < numKnots; i++ {
		knots[i] = newVector(0, 0)
	}

	tailPositions := util.NewSet[string]()
	tailPositions.Add(knots[len(knots) - 1].ToString())

	return &Rope{Knots: knots, TailPositions: tailPositions}
}

func (r *Rope) Move(movement *Vector) {
	destination := newVector(r.Knots[0].X + movement.X, r.Knots[0].Y + movement.Y)

	// While we're not at the destination, move the head
	for !r.Knots[0].IsEqual(destination) {
		headMovement := r.Knots[0].DirectionTo(destination)
		r.Knots[0].X += headMovement.X
		r.Knots[0].Y += headMovement.Y

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

func getDirection(str string) *Vector {
	switch str {
	case "R":
		return newVector(1, 0)
	case "U":
		return newVector(0, 1)
	case "D":
		return newVector(0, -1)
	case "L":
		return newVector(-1, 0)
	}

	panic("Invalid direction!")
}

func parseMovements(lines []string) []*Vector {
	movements := make([]*Vector, len(lines))

	for i, line := range lines {
		info := strings.Split(line, " ")
		direction := getDirection(info[0])
		distance := util.Atoi(info[1])

		movements[i] = newVector(direction.X * distance, direction.Y * distance)
	}

	return movements
}

func part1(movements []*Vector) {
	rope := newRope(2)

	for _, movement := range movements {
		rope.Move(movement)
	}

	fmt.Println(rope.TailPositions.Size())
}

func part2(movements []*Vector) {
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
