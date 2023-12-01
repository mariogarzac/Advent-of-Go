package main

import (
	"log"
	"strings"

    "github.com/mariogarzac/Advent/utils"
)

func GetMatchScore(filename string) int{

    sc, file, err := utils.OpenFile(filename)

    if err != nil { log.Fatal(err) }

    score := 0

    for sc.Scan(){
        match := strings.Fields(sc.Text())

        if sc.Text() != ""{
            switch strings.Join(match, " ") {
                //wins scissors rock, rock paper, paper scissors 
                case "A Y", "B Z", "C X": 
                score += 6
                case "A X", "B Y", "C Z": 
                score += 3
            default:
                score += 0
            }

            switch match[1] {
            case "X":
                score += 1
            case "Y":
                score += 2
            case "Z":
                score += 3
            default:
                // ignore the input
            }
        }
    }

    defer file.Close()

    return score
}

func GetStrategyScore(filename string) int {

    sc, file, err := utils.OpenFile(filename)

    if err != nil { log.Fatal(err) }

    score := 0

    for sc.Scan(){
        match := strings.Fields(sc.Text())

        if sc.Text() != ""{
            switch strings.Join(match, " ") {
            case "A X", "B X", "C X": // lose
                score += 0
            case "A Y", "B Y", "C Y": // ties
                score += 3
            case "A Z", "B Z", "C Z": // wins
                score += 6
            }

            switch strings.Join(match, " "){
            case "A X", "B Z", "C Y":
                score += 3
            case "A Z", "B Y", "C X":
                score += 2
            case "A Y", "B X", "C Z":
                score += 1
            default:
                // ignore the input
            }
        }
    }
    
    defer file.Close()
    return score
}

