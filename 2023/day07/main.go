package main

import (
    "fmt"
    "log"
    "sort"
    "strconv"
    "strings"

    "github.com/mariogarzac/Advent/utils"
)

type Hand struct {
    rank int
    cards string
    points int
}

type ByRankHighCard []Hand

func (a ByRankHighCard) Len() int      { return len(a) }
func (a ByRankHighCard) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRankHighCard) Less(i, j int) bool {
    if a[i].rank == a[j].rank {
        for k := 0; k < len(a[i].cards); k++ {
            c1 := convertions[string(a[i].cards[k])]
            c2 := convertions[string(a[j].cards[k])]

            if c1 != c2{
                return c1 < c2
            }
        }
    }

    return a[i].rank < a[j].rank
}

var convertions = map[string]int {
    "0" : 1,  "1" : 1,  "2" : 2, "3" : 3,
    "4" : 4,  "5" : 5,  "6" : 6, 
    "7" : 7,  "8" : 8,  "9" : 9,
    "T" : 10, "J" : 11, "Q" : 12,
    "K" : 13, "A" : 14,
}
var points = map[int]int{
    5: 6, 4: 5, 3: 3,
    2: 1, 1: 0,
}

func main() {
    file, err := utils.ReadWholeFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    games := parseCards(string(file))
    fmt.Println(part1(games))
    fmt.Println(part2(games))
}

func part1(games []Hand) int {

    for i := range games {
        games[i].rank = getHandRank(games[i].cards)
    }

    sort.Sort(ByRankHighCard(games))

    return calculatePoints(games)
}

func part2(games []Hand) int {

    for i := range games {
        jokerHand := jokerToCard(games[i].cards)
        games[i].rank = getHandRank(jokerHand)
        games[i].cards = strings.ReplaceAll(games[i].cards, "J", "0")
    }

    sort.Sort(ByRankHighCard(games))

    return calculatePoints(games)
}

func parseCards(game string) []Hand{
    hands := strings.Split(game, "\n")

    // remove last empty line
    hands = hands[0:len(hands) - 1]
    games := make([]Hand, len(hands))

    for i,h := range hands {
        hand := strings.Fields(h)

        points,_ := strconv.Atoi(hand[1])

        g := Hand{
            cards: hand[0],
            points: points,
        }

        games[i] =  g
    }
    return games
}

func getHandRank(hand string) int {
    cards, cardAmount := mapCards(hand)

    if cardAmount == 3 {
        if findPairs(cards) == 1{
            return 4
        }
    }

    if cardAmount == 2 {
        if findPairs(cards) == 2{
            return 2
        }
    }

    return points[cardAmount]
}

func mapCards(hand string) (map[int]int, int) {
    cards := map[int]int{}

    cardAmount := 0
    for _,c := range hand {
        card := convertions[string(c)]

        // initialize card
        if _, exists := cards[card]; !exists {
            cards[card] = 0
        }

        cards[card] += 1

        // keep track of highest play
        if cards[card] > cardAmount {
            cardAmount = cards[card]
        }
    }
    return cards, cardAmount
}

func jokerToCard(hand string) string {
    cards,_ := mapCards(hand)
    highest, card := 0, 0

    if cards[11] == 0{
        return hand
    }

    if cards[11] == 5{
        return "AAAAA"
    }

    for key,value := range cards{
        if key == 11{
            continue
        }
        if value > highest {
            card = key
            highest = value
        }
    }

    var targetCard string
    for key, c := range convertions {
        if c == card {
            targetCard = key
            break
        }
    }

    return strings.ReplaceAll(hand, "J", targetCard)
}

func findPairs(cards map[int]int) int{
    pairs := 0
    for _, value := range cards {
        if value == 2 {
            pairs += 1
        }
    }

    return pairs
}

func calculatePoints(games []Hand) int {
    total := 0
    bonus := 1
    for _, g := range games {
        total += (bonus * g.points)
        bonus += 1
    }

    return total
}
