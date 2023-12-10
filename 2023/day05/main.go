package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"

    "github.com/mariogarzac/Advent/utils"
)

type Block struct {
    destination int
    source int
    length int
}

func main() {

    file, err := utils.ReadWholeFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }


    fmt.Println(part1(string(file)))
}

func part1(line string) int {
    almanac := strings.Split(line, "\n\n")
    seeds := parseSeeds(almanac[0])
    blocks := almanac[1:]

    for _,block := range blocks {
        maps := [][]int{}
        splitRanges := parseBlock(block)
        maps = splitRanges

        newSeeds := []int{}
        for _,s := range seeds{
            found := false
            for _, m := range maps {
                srcStart, dstStart, rangeLen := m[0], m[1], m[2]
                if dstStart <= s && s < dstStart + rangeLen{
                    newSeeds = append(newSeeds, s - dstStart + srcStart)
                    found = true
                    break
                }
            }
            if !found {
                newSeeds = append(newSeeds, s)
            }
        }
        seeds = newSeeds
    }


    minSeed := seeds[0]
    for _, s := range seeds {
        if s < minSeed{
            minSeed = s
        }
    }
    
    return minSeed
}

func parseSeeds(s string) []int{
    str := strings.Split(s, " ")[1:]
    nums := make([]int,len(str))

    for i,n := range str{
        num,_ := strconv.Atoi(n)
        nums[i] = num

    }
    return nums
}

func parseBlock(block string) [][]int{ 

    blockRanges := strings.Split(block, "\n")[1:]

    numbers := [][]int{}

    // process each individual map
    for _,ranges := range blockRanges {

        // skip any empty ranges
        if ranges == ""{
            continue
        }

        // split the range by spaces
        n := []int{}
        blockRange := strings.Fields(ranges)

        // convert each number in the range to an int and append it
        for _,r := range blockRange {
            number,_ := strconv.Atoi(r)
            n = append(n, number)
        }

        // append the number slice to the slice that will be returned
        numbers = append(numbers, n)
    }

    return numbers
}
