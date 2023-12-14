package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/cznic/mathutil"

	"github.com/mariogarzac/Advent/utils"
)

type Directions struct {
    position string
    left string 
    right string
}

func main() {
    file, err := utils.ReadWholeFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    steps, mappy, count := parseInstructions(string(file))
    // fmt.Println(part1(steps, mappy))
    fmt.Println(part2(steps, mappy, count))

}

func parseInstructions(directions string) (string, map[string]Directions, []string) {

    lines := strings.Split(directions, "\n\n")
    steps := lines[0]

    mappy := map[string]Directions{}
    starts := []string{}

    network := strings.Split(lines[1], "\n")


    for _,line := range network {
        if line == ""{ continue }

        net := cleanLine(line)
        // remove double space and split
        n := strings.Split(strings.ReplaceAll(net, "  ", " "), " ")
        if n[0][len(n) - 1] == 'A' {
            starts = append(starts, n[0])
        }

        mappy[n[0]] = Directions{ left: n[1], right: n[2] }
    }

    return steps, mappy, starts
}

func part1(steps string, mappy map[string]Directions) int {

    dir := "AAA"
    i, count := 0,0
    size := len(steps)

    for dir != "ZZZ"{

        var newWay string
        if steps[i] == 'L' {
            newWay = mappy[dir].left
        }else{
            newWay = mappy[dir].right
        }

        dir = newWay

        i += 1
        count += 1

        if i == size {
            i = 0 
        }

        if dir == "ZZZ"{
            return count
        }
    }

    return count
}

func part2(steps string, network map[string]Directions, positions []string) int {

    cycles := [][]int{}
    for _, curr := range positions {
        cycle := []int{}

        firstZ := ""

        stepCount := 0
        i := 0

        for {
            for stepCount == 0 || curr[len(curr) - 1] != 'Z' {
                stepCount += 1
                if steps[i] == 'L' {
                    curr = network[curr].left
                }else{
                    curr = network[curr].right
                }

                i += 1

                if i == len(steps) { i = 0 }

            }

            cycle = append(cycle, stepCount)

            if firstZ == "" {
                firstZ = curr
            }else if firstZ == curr {
                break
            }
        }

        cycles = append(cycles, cycle)
    }

    lcm := cycles[0][0]

    for _, nums := range cycles[1:] {
        num := nums[0]
        lcm = (lcm * num) / int(mathutil.GCDUint64(uint64(lcm), uint64(num)))
    }

    return lcm
}

func cleanLine(line string) string{
    return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(line, "")
}
