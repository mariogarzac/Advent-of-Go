package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mariogarzac/Advent/tree"
	"github.com/mariogarzac/Advent/utils"
)

func FindMax(filename string) int {

    max, curr, num := 0,0,0
    var err error

    sc,file, err := utils.OpenFile(filename)

    if err != nil {log.Fatal(err)}
    

    for sc.Scan() {
        line := sc.Text()

        if line != "" {
            num, err = strconv.Atoi(line)
            if err != nil {log.Fatal(err)}
            curr += num
        }else{
            if curr > max {
                max = curr
            }
            curr = 0
        }
    }

    defer file.Close()
    return max
}

func FindTopThree(filename string) int {

    sc, file, err := utils.OpenFile(filename)
    bt := tree.BinaryTree{}
    num, curr := 0,0

    if err != nil { log.Fatal(err) }

    for sc.Scan(){
        line := sc.Text()

        if line != "" {
            num, err = strconv.Atoi(line)

            if err != nil {
                log.Fatal(err)
            }
            curr += num

        }else{
            bt.Insert(curr)
            curr = 0
        }
    }

    defer file.Close()

    return bt.FindTopThree(bt.Root)
}

func main(){
    fmt.Println(FindMax("input.txt"))
    fmt.Println(FindTopThree("input.txt"))
}
