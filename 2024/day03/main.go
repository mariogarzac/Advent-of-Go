package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/mariogarzac/Advent/utils"
)

func main() {
	sc, file, err := utils.OpenFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(part1(line))
		fmt.Println(part2(line))
	}

}

func part1(line string) int {
	re, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(line, -1)

	result := 0

	for _, match := range matches {
		digits := strings.Split(strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")"), ",")

		left, _ := strconv.Atoi(digits[0])
		right, _ := strconv.Atoi(digits[1])
		result += left * right
	}

	return result
}

func part2(line string) int {
	del, _ := regexp.Compile(`don't`)
	dontIndex := del.FindStringIndex(line)
	fmt.Println(line[dontIndex[0]:dontIndex[1]], dontIndex)

	do, _ := regexp.Compile(`do`)
	doIndex := do.FindStringIndex(line)
	fmt.Println(line[doIndex[0]:doIndex[1]], doIndex)

	re, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(line, 1)
	index := re.FindStringIndex(line)
	fmt.Println(line[index[0]:index[1]])

	fmt.Println(matches)
	result := 0

	return result
}
