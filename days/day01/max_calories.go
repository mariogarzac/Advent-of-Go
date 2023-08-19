package one

import (
	"bufio"
	"log"
	"strconv"
    "Advent/utils"
    "Advent/tree"
)

func FindMax(filename string) int {

    max, curr, num := 0,0,0
    var err error

    file, err := utils.OpenFile(filename)

    if err != nil {log.Fatal(err)}
    
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

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

    file, err := utils.OpenFile(filename)
    bt := tree.BinaryTree{}
    num, curr := 0,0

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    for scanner.Scan(){
        line := scanner.Text()

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
