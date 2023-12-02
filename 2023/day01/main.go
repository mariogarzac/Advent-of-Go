package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/mariogarzac/utils"
)

var numberMap = map[string]string {
    "one":   "one1one",
    "two":   "two2two",
    "three": "three3three",
    "four":  "four4four",
    "five":  "five5five",
    "six":   "six6six",
    "seven": "seven7seven",
    "eight": "eight8eight",
    "nine":  "nine9nine",
}

func main(){
    sc, file, err := utils.OpenFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    countP1, countP2 := 0,0

    for sc.Scan(){
        line := sc.Text()
        countP1 += part1(line)
        countP2 += part2(line)
    }

    fmt.Println(countP1)
    fmt.Println(countP2)
}

func part1(line string) int{

    var first,last byte
    var foundFirst bool

    foundFirst = false

    fmt.Println(line)

    // get all the numbers in the string 
    for i := range line{
        if line[i] >= '0' && line[i] <= '9'{
            if !foundFirst {
                first = line[i] - '0'
                foundFirst = true
            }
            last = line[i] - '0'
        }
    }

    fmt.Println(first, last)
    return int(first * 10 + last)
}

func part2(line string) int{

    for word, digit := range numberMap {
        if strings.Contains(line, word){
            line = strings.ReplaceAll(line, word, digit)
        }
    }

    number := part1(line)

    return number
}
