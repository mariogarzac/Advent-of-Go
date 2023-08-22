package main

import (
	"Advent/days/day01"
	"Advent/days/day02"
	"Advent/days/day03"
	"Advent/days/day04"
	"Advent/days/day05"
	"fmt"
)

func main(){

    // day one
    fmt.Println("Day one")
    fmt.Println(one.FindMax("day01/calories.txt"))
    fmt.Println(one.FindTopThree("day01/calories.txt"))
    
    fmt.Println()
    
    // day two
    fmt.Println("Day two")
    fmt.Println(two.GetMatchScore("day02/rps.txt"))
    fmt.Println(two.GetStrategyScore("day02/rps.txt"))
    
    fmt.Println()
    
    // day three
    fmt.Println("Day three")
    fmt.Println(three.RuckSack("day03/contents.txt"))
    fmt.Println(three.RuckSackPart2("day03/contents.txt"))
    
    fmt.Println()
    
    //day four
    fmt.Println("Day four")
    fmt.Println(four.Cleanup("day04/sections.txt"))
    fmt.Println(four.Cleanup2("day04/sections.txt"))
    
    //day five
    fmt.Println("Day five")
    fmt.Printf("%c ",five.OperateCrane("day05/boxes.txt"))
    fmt.Println()

}
