package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"
    "time"

    "github.com/mariogarzac/Advent/utils"
)

func main() {

    file, err := utils.ReadWholeFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    startNow := time.Now()
    fmt.Println(part1(string(file)))

    fmt.Println("Execution took", time.Since(startNow))

    startNow = time.Now()
    fmt.Println(part2(string(file)))
    fmt.Println("Execution took", time.Since(startNow))

}

func part1(line string) int {
    almanac := strings.Split(line, "\n\n")

    seeds := []int{}
    seeds = parseSeeds(almanac[0])

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

func part2(line string) int {

    seeds, maps := parseMaps(line)

    s := 0
    minSeed:= seeds[0]
    currSeed := minSeed

    // go through seed ranges
    for i := 1; i < len(seeds); i += 2 {
        // parse every seed in the given range
        for j := seeds[i - 1]; j < seeds[i - 1] + seeds[i]; j++{
            s = j
            // parse the seed through each individual map
            for _, m := range maps {
                for _, r := range m {
                    srcStart, dstStart, rangeLen := r[0], r[1], r[2]
                    if dstStart <= s && s < dstStart + rangeLen{
                        s = s - dstStart + srcStart
                        currSeed = s
                        break
                    }
                }
            }
            if currSeed < minSeed {
                minSeed = currSeed
            }
        }
    }

    return minSeed
}

func parseMaps(line string) ([]int, [][][]int){
    almanac := strings.Split(line, "\n\n")

    seeds := []int{}
    seeds = parseSeeds(almanac[0])

    maps := [][][]int{}
    blocks := almanac[1:]

    for _,block := range blocks {
        splitRanges := parseBlock(block)
        maps = append(maps, splitRanges)
    }

    return seeds, maps
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

func converToInt(strNums string) [][]int {

    splitNums := strings.Fields(strNums)
    nums := []int{}
    numbers := [][]int{}
    count := 0

    for _, num := range splitNums {
        n,err := strconv.Atoi(num)

        if err != nil {
            continue 
        }
        nums = append(nums, n)
        count += 1

        if count % 2 == 0{
            numbers = append(numbers, nums)
            nums = nil
        }
    }

    return numbers
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
