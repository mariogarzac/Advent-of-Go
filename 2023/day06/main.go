package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"

    "github.com/mariogarzac/Advent/utils"
)

func main() {
    file, err := utils.ReadWholeFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(part1(string(file)))
    fmt.Println(part2(string(file)))

}

func part1(line string) int {
    paper := strings.Split(line, "\n")
    time := convertToInt(strings.Fields(paper[0])[1:])
    record := convertToInt(strings.Fields(paper[1])[1:])

    total := 1
    for i := 0; i < len(record); i++{
        count := 0
        for j := 0; j < time[i]; j++{
            if j * (time[i] - j) > record[i]{
                count += 1
            }
        }
        total *= (count)
    }

    return total
}

func part2(line string) int {
    paper := strings.Split(line, "\n")
    time := joinNums(strings.Fields(paper[0])[1:])
    record := joinNums(strings.Fields(paper[1])[1:])

    count := 0
    for i := 0; i < time; i++{
        if i * (time - i) > record{
            count += 1
        }
    }

    return count
}

func joinNums(strNums []string) int{
    var sb strings.Builder

    for _, n := range strNums {
        sb.WriteString(n)
    }

    num, _ := strconv.Atoi(sb.String())
    return num
}

func convertToInt(strNums []string) []int {
    nums := make([]int, len(strNums))
    for i, n := range strNums {
        num, _ := strconv.Atoi(n)
        nums[i] = num
    }

    return nums
}
