package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
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
    almanac := strings.Split(line, "\n\n")

    seeds := converToInt(almanac[0])
    blocks := almanac[1:]

    ch := make(chan int)
    var wg sync.WaitGroup

    locations := []int{}
    for _, s := range seeds {
        wg.Add(1)
        go concurrentP2(s, blocks, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for result := range ch {
        locations = append(locations, result)
    }

    minSeed := locations[0]
    for _, s := range locations {
        if s < minSeed {
            minSeed = s
        }
    }

    return minSeed
}
func concurrentP2(seeds []int, blocks []string, ch chan<-int, wg *sync.WaitGroup) {

    defer wg.Done()

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

    ch <- minSeed
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
