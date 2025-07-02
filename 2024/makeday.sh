#!/bin/sh

mkdir day0"$1"
cd day0"$1"

touch main.go input.txt test.txt

echo "
package main

import (
	\"fmt\"
	\"log\"

	\"github.com/mariogarzac/Advent/utils\"
)

func main() {
	sc, file, err := utils.OpenFile(\"test.txt\")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for sc.Scan() {
		line := sc.Text()
		part1(line)
	}

}

func part1(line string) int {
	fmt.Println(line)

	return 0
}
" > main.go
