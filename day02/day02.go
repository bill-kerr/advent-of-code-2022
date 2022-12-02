package day02

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

type Result int

const (
	Lost Result = iota
	Draw
	Won
)

var scores map[rune]int = map[rune]int{
	'A': 1,
	'B': 2,
	'C': 3,
}

var choiceMap map[rune]rune = map[rune]rune {
	'X': 'A',
	'Y': 'B',
	'Z': 'C',
}

var resultMap map[rune]Result = map[rune]Result {
	'X': Lost,
	'Y': Draw,
	'Z': Won,
}

func calculateScore(choice rune, result Result) int {
	score := scores[choice]

	if result == Lost {
		return score
	} else if result == Draw {
		return score + 3
	}

	return score + 6
}

// func getResult(myChoice rune, opponentChoice rune) Result {
// 	convertedChoice := choiceMap[myChoice]

// 	if opponentChoice == convertedChoice {
// 		return Draw
// 	}

// 	if opponentChoice == 'A' && convertedChoice == 'C' {
// 		return Lost
// 	}

// 	if opponentChoice == 'A' && convertedChoice == 'B' {
// 		return Won
// 	}

// 	if opponentChoice == 'B' && convertedChoice == 'A' {
// 		return Lost
// 	}

// 	if opponentChoice == 'B' && convertedChoice == 'C' {
// 		return Won
// 	}

// 	if opponentChoice == 'C' && convertedChoice == 'A' {
// 		return Won
// 	}

// 	return Lost
// }

func getRequiredPlay(opponentChoice rune, requiredResult Result) rune {
	if requiredResult == Lost {
		if opponentChoice == 'A' {
			return 'C'
		}

		if opponentChoice == 'B' {
			return 'A'
		}

		return 'B'
	}

	if requiredResult == Draw {
		return opponentChoice
	}

	if requiredResult == Won {
		if opponentChoice == 'A' {
			return 'B'
		}

		if opponentChoice == 'B' {
			return 'C'
		}

		return 'A'
	}

	return 'A'
}

func Run() {
	lines := util.OpenAndRead("./day02/input.txt");
	score := 0

	for _, val := range lines {
		values := []rune(val)

		opponentChoice := values[0]
		result := resultMap[values[2]]
		myChoice := getRequiredPlay(opponentChoice, result)

		// result := getResult(myChoice, opponentChoice)
		roundScore := calculateScore(myChoice, result)

		fmt.Println(string(opponentChoice), string(choiceMap[myChoice]), result, roundScore)

		score += roundScore
	}

	fmt.Println(score)
}