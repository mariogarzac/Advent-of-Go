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

func main() {
	sc, file, err := utils.OpenFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	warehouse := make([][]string, 0)
	i := 0
	ipos, jpos := 0, 0
	for sc.Scan() {
		line := sc.Text()
		aisle := strings.Split(line, "")
		warehouse = append(warehouse, aisle)
		for j, char := range aisle {
			if char == "^" {
				ipos, jpos = i, j
			}
		}
		i += 1
	}
	directions := []Point{
		{-1, 0}, // north
		{0, 1},  // east
		{1, 0},  // south
		{0, -1}, // west
	}

	p1, total, exit := 0, 0, false
	visited := make(map[Point]bool)
	for !exit {
		for _, dir := range directions {
			total, exit = part1(warehouse, &ipos, &jpos, dir, visited)
			p1 += total
			if exit {
				break
			}
		}
	}
	fmt.Println("total: ", p1)
}

func part1(warehouse [][]string, i, j *int, dir Point, visited map[Point]bool) (int, bool) {
	exit, steps := false, 0

	for {
		ni, nj := *i+dir.i, *j+dir.j
		if ni >= len(warehouse) || ni < 0 || nj >= len(warehouse[0]) || nj < 0 {
			exit = true
			return steps, exit
		}

		if warehouse[ni][nj] == "#" {
			break
		}

		*i, *j = ni, nj

		if !visited[Point{*i, *j}] {
			visited[Point{*i, *j}] = true
			steps += 1
		}
	}
	return steps, exit
}
