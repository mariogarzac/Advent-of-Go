package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mariogarzac/Advent/utils"
)

func main() {
	sc, file, err := utils.OpenFile("day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	p1, p2 := 0, 0

	for sc.Scan() {
		line := sc.Text()
		report := make([]int, 0)

		stringReport := strings.Fields(line)

		for _, item := range stringReport {
			num, _ := strconv.Atoi(item)

			report = append(report, num)
		}

		p1 += part1(report)
		p2 += part2(report)
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func part1(line []int) int {
	return testReport(line)
}

func part2(line []int) int {
	if testReport(line) > 0 {
		return 1
	} else {
		return testLevel(line)
	}
}

func testReport(level []int) int {
	increasing := level[1] > level[0]
	for i := 1; i < len(level); i++ {
		diff := utils.Abs(level[i] - level[i-1])
		validDiff := 1 <= diff && diff <= 3
		isSequential := (increasing && level[i] > level[i-1]) || (!increasing && level[i] < level[i-1])

		if !validDiff || !isSequential {
			return 0
		}
	}
	return 1
}

func testLevel(level []int) int {
	for i := 0; i < len(level); i++ {
		if testReport(deleteLevelAt(i, level)) > 0 {
			return 1
		}
	}
	return 0
}

func deleteLevelAt(idx int, levels []int) []int {
	deleted := make([]int, len(levels)-1)
	copy(deleted[:idx], levels[:idx])
	copy(deleted[idx:], levels[idx+1:])
	return deleted
}
