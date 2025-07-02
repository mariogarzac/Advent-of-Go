package main

import (
	"fmt"
	"log"
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

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(part1(line))
		fmt.Println(part2(line))
	}

}

func part1(line string) int {
	total := 0
	for i := 0; i < len(line); i++ {
		if line[i] == 'm' && line[i+3] == '(' {
			exs, finish := isMul(i+4, line)
			if exs {
				first, second := findNumbers(line[i+4 : finish])
				total += second * first
			}
		}
	}

	return total
}

func part2(line string) int {
	total := 0
	shouldDo := true

	for i := 0; i < len(line); i++ {
		if line[i] == 'd' && line[i+2] == 'n' {
			shouldDo = false
		} else if line[i] == 'd' && line[i+1] == 'o' {
			shouldDo = true
		}

		if line[i] == 'm' && line[i+3] == '(' {
			exs, finish := isMul(i+4, line)
			if exs && shouldDo {
				first, second := findNumbers(line[i+4 : finish])
				total += second * first
			}
		}
	}

	return total
}

func findNumbers(line string) (int, int) {
	nums := strings.Split(line, ",")
	first, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0
	}
	second, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0
	}

	return first, second
}

func isMul(index int, line string) (bool, int) {
	for i := index; i < index+9; i++ {
		if line[i] == ')' {
			return true, i
		}
	}
	return false, 0
}
