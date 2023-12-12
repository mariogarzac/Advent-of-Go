package main

import (
    "fmt"
    "log"
    "regexp"
    "strings"

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

    steps, mappy := parseInstructions(string(file))
    fmt.Println(part1(steps, mappy))

}

func parseInstructions(directions string) (string, map[string]Directions) {

    lines := strings.Split(directions, "\n\n")
    steps := lines[0]

    mappy := map[string]Directions{}

    network := strings.Split(lines[1], "\n")


    for _,line := range network {
        if line == ""{ continue }

        net := cleanLine(line)
        // remove double space and split
        n := strings.Split(strings.ReplaceAll(net, "  ", " "), " ")

        mappy[n[0]] = Directions{ left: n[1], right: n[2] }
    }

    return steps, mappy
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

func cleanLine(line string) string{
    return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(line, "")
}

