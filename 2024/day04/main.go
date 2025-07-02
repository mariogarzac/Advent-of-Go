package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mariogarzac/Advent/utils"
)

type Point struct {
	i int
	j int
}

var crossword = make([][]string, 0)

func main() {
	sc, file, err := utils.OpenFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for sc.Scan() {
		line := sc.Text()
		crossword = append(crossword, strings.Split(line, ""))
	}
	fmt.Println(part1() / 2)
	fmt.Println(part2())
}

func part1() int {
	total := 0
	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for i := range crossword {
		for j := range crossword[i] {
			letter := crossword[i][j]
			if letter == "X" || letter == "S" {
				for _, dir := range directions {
					total += traverse(0, i, j, dir, crossword[i][j])
				}
			}
		}
	}
	return total
}

func part2() int {
	total := 0
	leftDiagonal := []Point{
		{-1, -1}, {1, 1},
	}
	rightDiagonal := []Point{
		{-1, 1}, {1, -1},
	}
	for i := range crossword {
		for j := range crossword[i] {
			letter := crossword[i][j]
			if letter == "A" {
				if searchForX(i, j, leftDiagonal, rightDiagonal) > 0 {
					total += 1
				}
			}
		}
	}
	return total
}

func searchForX(istart, jstart int, topLeft, topRight []Point) int {

	for _, dir := range topLeft {
		idx, idj := istart+dir.i, jstart+dir.j
		if idj < 0 || idx < 0 || idx >= len(crossword) || idj >= len(crossword[0]) {
			return 0
		}
	}

	left := crossword[istart+topLeft[0].i][jstart+topLeft[0].j] + "A" + crossword[istart+topLeft[1].i][jstart+topLeft[1].j]

	for _, dir := range topRight {
		idx := istart + dir.i
		idj := jstart + dir.j
		if idj < 0 || idx < 0 || idx >= len(crossword) || idj >= len(crossword[0]) {
			return 0
		}
		right := crossword[istart+topRight[0].i][jstart+topRight[0].j] + "A" + crossword[istart+topRight[1].i][jstart+topRight[1].j]

		if hasSubstringMAS(left) && hasSubstringMAS(right) {
			return 1
		}
	}
	return 0
}

func traverse(total, idx, idj int, dir Point, substr string) int {
	i, j := idx+dir.i, idj+dir.j
	if i < 0 || j < 0 || i >= len(crossword) || j >= len(crossword[0]) || len(substr) > 4 {
		return total
	}
	substr += string(crossword[i][j])
	if hasSubstringXMAS(substr) {
		if len(substr) == 4 && (substr == "XMAS" || substr == "SAMX") {
			total++
		}
		total = traverse(total, i, j, dir, substr)
	}
	return total
}

func hasSubstringXMAS(substr string) bool {
	return strings.HasPrefix("XMAS", substr) || strings.HasPrefix("SAMX", substr)
}

func hasSubstringMAS(substr string) bool {
	return strings.HasPrefix("MAS", substr) || strings.HasPrefix("SAM", substr)
}
