package main

import (
	"log"

    "github.com/mariogarzac/Advent/utils"
)

func Solve(filename string) {
    one := FindSignal(filename, 1)
    two := FindSignal(filename, 2)
    print(one, "\n")
    print(two, "\n")
}

func isUnique(stream []byte) bool {

    exists := make(map[byte]bool)

    for _, char := range stream {

        if exists[char]{
            return false
        }

        exists[char] = true
    }

    return true
}

func FindSignal(filename string, part int) int {

    window := 4
    if part == 2 { window = 14 }

    signal,err := utils.ReadWholeFile(filename)
    if err != nil { log.Fatal(err) }

    var left int

    for i := window; i < len(signal); i++{
        left = i - window
        if isUnique(signal[left:i]){
           return i
        }
    }

    return 0

}

