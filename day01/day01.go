package day01

import (
	"log"
	"sort"
	"strconv"

	"github.com/bill-kerr/advent-of-code-2022/util"
)

func Run() {
	lines := util.OpenAndRead("./day01/input.txt")

	calorieCounts := []int{}

	elfNum := 0

	max := 0

	for _, val := range lines {
		if val == "" {
			elfNum++
			continue
		}

		calories, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		if elfNum >= len(calorieCounts) {
			calorieCounts = append(calorieCounts, 0)
		}

		calorieCounts[elfNum] += calories

		if calorieCounts[elfNum] > max {
			max = calorieCounts[elfNum]
		}
	}

	sort.Ints(calorieCounts)
	
	topThree := calorieCounts[len(calorieCounts) - 3:]

	log.Println(max, util.SumSlice(topThree))
}