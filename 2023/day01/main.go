package main

import (
    "fmt"
    "strconv"

    "github.com/mariogarzac/utils"
)

func part1() int{
    sc, file, err := utils.OpenFile("day01/input.txt")

    if err != nil {
        return 0
    }

    defer file.Close()

    var nums []byte
    count := 0

    for sc.Scan(){
        line := sc.Text()

        // get all the numbers in the string 
        for i := range line{
            if line[i] >= '0' && line[i] <= '9'{
                nums = append(nums, line[i] - '0')
            }

        }

        // get the first and the last byte from the slice and convert them to a string
        num := strconv.Itoa(int(nums[0])) + strconv.Itoa(int(nums[len(nums) - 1 ]))

        // convert the string back to an int lmao
        result,_ := strconv.Atoi(num)

        // add the number and reset the array
        count += result
        nums = nums[:0]
    }

    return count
}

func main(){
    fmt.Println(part1())
}

