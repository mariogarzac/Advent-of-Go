package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/mariogarzac/Advent/utils"
)

type Sack struct {
    red   int
    green int
    blue  int
}

var s = Sack{
    red: 12,
    green: 13,
    blue: 14,
}


func main(){
    sc, file, err := utils.OpenFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    gamesP1, gamesP2 := 0,0
    var p2Sack Sack

    for sc.Scan(){
        line := sc.Text()

        gamesP1 += part1(line)

        p2Sack = part2(line)
        gamesP2 += (p2Sack.red * p2Sack.green * p2Sack.blue)

    }

    fmt.Println(gamesP1)
    fmt.Println(gamesP2)
}

func part1(line string) int{

    line = cleanLine(line)
    gameSets := strings.Split(line, " ")

    gameID := 0

    for i, game := range gameSets {
        switch game{
        case "Game":
            gameID,_ = strconv.Atoi(gameSets[i + 1]) 

        case "red":
            cubes,_ := strconv.Atoi(gameSets[i - 1])
            if cubes > s.red{
                gameID = 0
            }

        case "green":
            cubes,_ := strconv.Atoi(gameSets[i - 1])
            if cubes > s.green{
                gameID = 0
            }
            
        case "blue":
            cubes,_ := strconv.Atoi(gameSets[i - 1])
            if cubes > s.blue{
                gameID = 0
            }
            
        default:
            continue
        }
    }

    return gameID
}

func part2(line string) Sack{
    line = cleanLine(line)
    gameSets := strings.Split(line, " ")

    sack := Sack{
        red:   0,
        green: 0,
        blue:  0,
    }

    for i, game := range gameSets {
        switch game{
        case "red":
            cubes,_ := strconv.Atoi(gameSets[i - 1])
            if cubes > sack.red{
                sack.red = cubes
            }

        case "green":
            cubes,_ := strconv.Atoi(gameSets[i - 1])
            if cubes > sack.green{
                sack.green = cubes
            }

        case "blue":
            cubes,_ := strconv.Atoi(gameSets[i - 1])
            if cubes > sack.blue{
                sack.blue = cubes
            }

        default:
            continue
        }
    }

    return sack

}

func cleanLine(line string) string{
     return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(line, "")
}

