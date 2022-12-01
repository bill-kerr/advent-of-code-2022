package util

import (
	"bufio"
	"log"
	"os"
)

func OpenAndRead(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to read text file")
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	
	lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func SumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}