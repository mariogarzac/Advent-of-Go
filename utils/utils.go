package utils

import (
	"bufio"
	"log"
	"os"
)

func OpenFile(filename string) (*bufio.Scanner, *os.File, error) {

    patherino := "../Advent/days/"
    file, err := os.Open(patherino + filename)

    if err != nil {
        log.Fatal(err)
        return nil, file, err
    }

    sc := bufio.NewScanner(file)
    return sc, file, nil
}

type Stack []interface{}

func (s *Stack) Push(item interface{}){
    *s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, bool){

    if len(*s) == 0{
        return 0, false
    }

    index := len(*s) - 1
    item := (*s)[index]
    *s = (*s)[:index]

    return item, true
}

func (s *Stack) Top() (interface{}){

    if len(*s) == 0{
        return 0
    }

    index := len(*s) - 1
    item := (*s)[index]

    return item
}

func (s *Stack) Insert(index int, item interface{}) {
    if index < 0 || index > len(*s) {
        // Invalid index
        return
    }

    *s = append(*s, nil) // Expand the slice
    copy((*s)[index+1:], (*s)[index:]) // Shift elements to the right
    (*s)[index] = item // Insert the item at the specified index
}
