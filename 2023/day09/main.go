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

    countP1 := 0
    countP2 := 0

    for sc.Scan(){
        line := sc.Text()
        countP1 += part1(line)
        countP2 += part2(line)
    }

    fmt.Println(countP1)
    fmt.Println(countP2)

}

func getReport(line string) [][]int {
    report := [][]int{}
    report = append(report, getInts(strings.Fields(line))) 
    index := 0

    for !checkZeros(report[len(report) - 1]) {
        first := report[index][0]
        newValues := []int{}

        for _,num := range report[index][1:] {
            newValues = append(newValues, num - first)
            first = num
        }
        report = append(report, newValues)
        index += 1
    }

    return report
}

func part1(line string) int {

    report := getReport(line)

    return extrapolateForward(report)
}

func part2(line string) int {

    report := getReport(line)

    return extrapolateBackward(report)
}

func extrapolateForward(report [][]int) int {
    index := len(report) - 2
    numToAdd := report[index][len(report[index]) - 1]

    for i := index - 1; i >= 0; i -- {
        nums := report[i]
        numToAdd = nums[len(nums) - 1] + numToAdd
    }

    return numToAdd
}

func extrapolateBackward(report [][]int) int {
    index := len(report) - 2

    numToAdd := report[index][0]

    for i := index - 1; i >= 0; i -- {
        nums := report[i]
        numToAdd = nums[0] + (numToAdd * -1)
    }

    return numToAdd
}

func checkZeros(nums []int) bool {

    for _,n := range nums {
        if n != 0 {
            return false
        }
    }

    return true
}

func getInts(line []string) []int {

    nums := make([]int, len(line))
    for i, value := range line {
        n,_ := strconv.Atoi(value)
        nums[i] = n

    }

    return nums
}
