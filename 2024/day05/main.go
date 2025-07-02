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

	orderMap := make(map[int]map[int]bool)
	updates := make([][]int, 0)
	next := false

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			next = true
			continue
		}

		if !next {
			pageOrders := strings.Split(line, "|")
			x, _ := strconv.Atoi(pageOrders[0])
			y, _ := strconv.Atoi(pageOrders[1])

			if _, exists := orderMap[x]; !exists {
				orderMap[x] = make(map[int]bool)
			}

			orderMap[x][y] = true

		} else {
			update := make([]int, 0)
			test := strings.Split(line, ",")
			for _, digits := range test {
				num, _ := strconv.Atoi(digits)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}
	totalP1, totalP2 := 0, 0
	fails := make([][]int, 0)

	for _, update := range updates {
		totalP1 += part1(orderMap, update, &fails)
	}

	for _, fail := range fails {
		totalP2 += part2(orderMap, fail)
	}

	// for key, value := range orderMap {
	// 	fmt.Println(key, value)
	// }

	fmt.Println(totalP1)
	fmt.Println(totalP2)
}

func part1(orderMap map[int]map[int]bool, update []int, fails *[][]int) int {
	for i := 0; i < len(update); i++ {
		checking := update[i]
		for j := i + 1; j < len(update); j++ {
			num := update[j]
			if _, exists := orderMap[checking][num]; !exists {
				*fails = append(*fails, update)
				return 0
			}
		}
	}
	return update[len(update)/2]
}

func part2(orderMap map[int]map[int]bool, fail []int) int {
	for i := 0; i < len(fail); i++ {
		for j := i + 1; j < len(fail); j++ {
			checking, num := fail[i], fail[j]
			if _, exists := orderMap[checking][num]; !exists {
				fail[i], fail[j] = fail[j], fail[i]
			}
		}
	}
	return fail[len(fail)/2]
}
