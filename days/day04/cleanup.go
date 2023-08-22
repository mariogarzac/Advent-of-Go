package four

import (
	"Advent/utils"
	"log"
	"strconv"
	"strings"
)

func turnToInts(s []string) []int {

    var sec []int
    for _, i := range s {
        j, err := strconv.Atoi(i)

        if err != nil { panic(err) }

        sec = append(sec, j)
    }

    return sec
}

// Checks if the whole range is included
func isInRange(s []int) bool {

    // check if first pair includes second
    if s[0] <= s[2] && s[1] >= s[3] {
        return true
    }

    // check if second pair includes first
    if s[2] <= s[0] && s[3] >= s[1] {
        return true
    }

    return false
}

func Cleanup(filename string) int{

    sc, file, err := utils.OpenFile(filename)

    if err != nil { log.Fatal(err) }

    var sections []string
    var ints []int
    counter := 0

    for sc.Scan(){
        sections = strings.FieldsFunc(sc.Text(), func(r rune) bool { 
            return r == ',' || r == '-' 
        })

        ints = turnToInts(sections)
        if isInRange(ints) {
            counter++
        }
    }

    defer file.Close()
    return counter
}

// Checks if at least one part overlaps
func overlaps(range1, range2 []int) bool {

    return (range1[0] <= range2[0] && range2[0] <= range1[1]) ||
           (range2[0] <= range1[0] && range1[0] <= range2[1])
}

func Cleanup2(filename string) int{

    sc, file, err := utils.OpenFile(filename)

    if err != nil { log.Fatal(err) }

    var sections []string
    var ints []int
    counter := 0

    for sc.Scan(){

        sections = strings.FieldsFunc(sc.Text(), func(r rune) bool { 
            return r == ',' || r == '-' 
        })

        ints = turnToInts(sections)

        if overlaps(ints[:2], ints[2:]) {
            counter++
        }
    }

    defer file.Close()
    return counter
}
