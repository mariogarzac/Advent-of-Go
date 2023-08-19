package main

import (
    "Advent/days/day01"
    "Advent/days/day02"
    "Advent/days/day03"
    "Advent/days/day04"
    "fmt"
)

func main(){

    // day one
    fmt.Println("Day one")
    fmt.Println(one.FindMax("days/days01/calories.txt"))
    fmt.Println(one.FindTopThree("days/days01/calories.txt"))
    
    fmt.Println()
    
    // day two
    fmt.Println("Day two")
    fmt.Println(two.GetMatchScore("days/days02/rps.txt"))
    fmt.Println(two.GetStrategyScore("days/days02/rps.txt"))
    
    fmt.Println()

    // day three
    fmt.Println("Day three")
    fmt.Println(three.RuckSack("days/days03/contents.txt"))
    fmt.Println(three.RuckSackPart2("days/days03/contents.txt"))
    
    fmt.Println()

    //day four
    fmt.Println("Day four")
    fmt.Println(four.Cleanup("days/days04/sections.txt"))
    fmt.Println(four.Cleanup2("days/days04/sections.txt"))
}
