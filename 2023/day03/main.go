package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/mariogarzac/Advent/utils"
)

type Point struct {
    X int
    Y int
}

func (p Point) addPoint(d Point) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func main(){
    b,err := utils.ReadWholeFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }


    symbols := getSymbols(b)
    engineParts := getEngineParts(b, symbols)

    fmt.Println(part1(engineParts))
    fmt.Println(part2(symbols, engineParts))

}

func getSymbols(file []byte) map[Point]string {

    symbols := map[Point]string{}

    for y, sym := range strings.Fields(string(file)) {
        for x, r := range sym {
            if r != '.' && !unicode.IsDigit(r){
                symbols[Point{x,y}] = string(r)
            }
        }
    }

    return symbols
}

func getEngineParts(file []byte, symbols map[Point]string) map[Point][]int {
    engineParts := map[Point][]int{}
    re := regexp.MustCompile(`\d+`)

    directions := []Point{
        {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
    }

    for y, s := range strings.Fields(string(file)){
        for _, match := range re.FindAllStringIndex(s, -1) {
            keys := map[Point]bool{}

            for x := match[0]; x < match[1]; x++ {
                for _, d := range directions {
                    keys[Point{x,y}.addPoint(d)] = true
                }
            }

            n, _ := strconv.Atoi(s[match[0]:match[1]])
            for p := range keys {
                if _, exists := symbols[p]; exists {
                    engineParts[p] = append(engineParts[p], n)
                }
            }
        }
    }

    return engineParts
}

func part1(engineParts map[Point][]int) int {

    partNumbers := 0
    for _, values := range engineParts {
        for _, value := range values {
            partNumbers += value
        }
    }
    return partNumbers
}

func part2(symbols map[Point]string, engineParts map[Point][]int) int {

    partNumbers := 0
    ratio := 1

    for keys, values := range engineParts {
        symbol := symbols[keys]

        if len(values) > 1 && symbol == "*" {
            for _, value := range values {
                ratio *= value
            }

            partNumbers += ratio
            ratio = 1
        }
    }

    return partNumbers
}
