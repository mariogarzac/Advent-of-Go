package main

import (
    "Advent/days/one"
    "Advent/days/two"
    "Advent/days/three"
    "Advent/days/four"
    "fmt"
)

func main(){

    // day one
    fmt.Println("Day one")
    fmt.Println(one.FindMax("days/one/calories.txt"))
    fmt.Println(one.FindTopThree("days/one/calories.txt"))
    
    fmt.Println()
    
    // day two
    fmt.Println("Day two")
    fmt.Println(two.GetMatchScore("days/two/rps.txt"))
    fmt.Println(two.GetStrategyScore("days/two/rps.txt"))
    
    fmt.Println()

    // day three
    fmt.Println("Day three")
    fmt.Println(three.RuckSack("days/three/contents.txt"))
    fmt.Println(three.RuckSackPart2("days/three/contents.txt"))
    
    fmt.Println()

    //day four
    fmt.Println("Day four")
    fmt.Println(four.Cleanup("days/four/sections.txt"))
}
