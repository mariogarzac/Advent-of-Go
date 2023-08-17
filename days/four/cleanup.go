package four

import (
	"Advent/utils"
	"bufio"
	"log"
	"strconv"
	"strings"
)

// Checks if the whole range is included
func isInRange(s []string) bool {

    var sec []int

    for _, i := range s {
        j, err := strconv.Atoi(i)

        if err != nil { panic(err) }

        sec = append(sec, j)
    }

    // check if first pair includes second
    if sec[0] <= sec[2] && sec[1] >= sec[3] {
        return true
    }

    // check if second pair includes first
    if sec[2] <= sec[0] && sec[3] >= sec[1] {
        return true
    }

    return false
}

// Checks if at least one part overlaps
func isPair(s []string) bool {

    return false
}

func Cleanup(filename string) int{

    file, err := utils.OpenFile(filename)

    if err != nil { log.Fatal(err) }

    sc := bufio.NewScanner(file)

    var sections []string
    var tmp []string
    counter := 0

    for sc.Scan(){
       tmp = strings.Split(sc.Text(), ",")
       sections = strings.Split(strings.Join(tmp, " "), "-")
       sections = strings.Split(strings.Join(sections, " "), " ")

       if isInRange(sections) {
           counter++
       }
    }
    return counter
}

