package main

import (
    "fmt"
    "log"
    "math"
    "strconv"
    "strings"

    "github.com/mariogarzac/Advent/utils"
)

func main() {
    sc, file, err := utils.OpenFile("input.txt")
    totalCards := make(map[int]int,6)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    countP1 := 0

    i := 1
    for sc.Scan(){
        line := sc.Text()
        points := part1(line)


        countP1 += calculate(points)
        totalCards[i] = points
        i += 1
    }

    countP2 := part2(totalCards)
    fmt.Println(countP1)
    fmt.Println(countP2)
}

func part1(line string) int {

    nums := strings.Split(line, ":")[1:]
    count := 0

    winningNumbers := map[int]bool{}

    match := strings.Split(strings.Join(nums, " "), "|")

    winningNums := strings.Fields(match[0])
    elfNumbers := strings.Fields(match[1])

    for _,strNum := range winningNums{
        num,_ := strconv.Atoi(strNum)
        winningNumbers[num] = true
    }

    for _,strNum := range elfNumbers{
        num,_ := strconv.Atoi(strNum)

        if winningNumbers[num] {
            count += 1
        }
    }
    return count
}

type Card struct {
    id int
    matchingCards int
    wonCards int
    children []int
}


func part2(cards map[int]int) int { 

    cardList := make([]*Card, len(cards) + 1)

    // add the cards to the map 
    for i, matches := range cards {
        points := 0
        if matches == 0{
            points = 1
        }
        cardList[i] = &Card{id: i, matchingCards : matches, wonCards: points, children : make([]int, matches)}

        // add the additional cards to each card
        curr := cardList[i]

        for j := i + 1; j <= curr.matchingCards + i; j++ {
           curr.children[j - i - 1] = j
        }
    }

    count := 0

    // start at the end
    for i := len(cardList) - 1; i > 0; i-- {
        size := len(cardList[i].children)
        currCard := cardList[i]

        if size == 0 {
            count += 1
            continue
        }

        for j := 0; j < size; j++ {
            child := currCard.children[j]

            currCard.wonCards += cardList[child].wonCards
        }

        currCard.wonCards +=  1
        count += currCard.wonCards
    }


    return count
}

func calculate(p int) int { 

    if p == 0{
        return 0
    }
    if p == 1 {
        return 1
    }
    return int(math.Pow(2,float64(p-1)))
}

