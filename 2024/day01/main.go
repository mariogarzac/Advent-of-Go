package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/mariogarzac/Advent/utils"
)

func main() {
	sc, file, err := utils.OpenFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	firstCol := make([]int, 0)
	secondCol := make([]int, 0)

	for sc.Scan() {

		line := sc.Text()

		splitLine := strings.Fields(line)

		firstColNum, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}

		secondColNum, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}

		firstCol = append(firstCol, firstColNum)
		secondCol = append(secondCol, secondColNum)
	}

	fmt.Println(part1(firstCol, secondCol))
	fmt.Println(part2(firstCol, secondCol))
}

func part1(first []int, second []int) int {
	totalDistance := 0

	slices.Sort(first)
	slices.Sort(second)

	for i, num := range first {
		if num > second[i] {
			totalDistance += num - second[i]
		} else {
			totalDistance += second[i] - num
		}
	}

	return totalDistance
}

func part2(first []int, second []int) int {
	similarityList := make(map[int]int)
	similarityScore := 0

	for _, item := range first {
		if _, exists := similarityList[item]; !exists {
			similarityList[item] = 0
		}
	}

	for _, item := range second {
		if _, exists := similarityList[item]; exists {
			similarityList[item] += 1
		}
	}

	for number, repeats := range similarityList {
		similarityScore += number * repeats
	}

	return similarityScore
}
