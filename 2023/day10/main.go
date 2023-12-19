package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/mariogarzac/Advent/utils"
)

type Point struct {
    x int
    y int
}

var valid  = map[Point]map[string]map[string]bool  {
    {0, 1} : {
        // - J 7
        "S" : {"-": true, "J": true, "7": true}, 
        "-" : {"-": true, "J": true, "7": true},
        "L" : {"-": true, "J": true, "7": true},
        "F" : {"-": true, "J": true, "7": true},
    },

    {0,-1} : {
        // - F L
        "S" : {"-": true, "F": true, "L": true}, 
        "-" : {"-": true, "F": true, "L": true},
        "J" : {"-": true, "F": true, "L": true},
        "7" : {"-": true, "F": true, "L": true},
    },

    {-1,0} : {
        // | F 7
        "S" : {"|": true, "F": true, "7": true}, 
        "|" : {"|": true, "F": true, "7": true}, 
        "L" : {"|": true, "F": true, "7": true},
        "J" : {"|": true, "F": true, "7": true},
    },

    {1, 0} : {
        // | L J
        "S" : {"|": true, "L": true, "J": true,}, 
        "|" : {"|": true, "L": true, "J": true,}, 
        "7" : {"|": true, "L": true, "J": true,},
        "F" : {"|": true, "L": true, "J": true,},
    },
}

var directions = []Point { 
    {0, 1}, {0, -1}, {-1, 0}, {1, 0},
}


func main() {
    file, err := utils.ReadWholeFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    var pipes[][]string
    var row, col int
    row, col, pipes = parseTubes(string(file))

    fmt.Println(part1(row, col, pipes))

}

func part1(row, col int, pipes [][]string) int {

    begin := getPossiblePaths(row, col, pipes)
    paths := [][]Point{}

    for _, b := range begin {
        visited := make(map[Point]bool)
        path := []Point{}
        traverseMaze(b.x, b.y, pipes, visited, &path)
        paths = append(paths, path)
    }

    return findCommonPosition(paths)
}

func replaceS(){
}

func traverseMaze(row, col int, pipes [][]string, visited map[Point]bool, path *[]Point) {
    visited[Point{row, col}] = true

    *path = append(*path, Point{row,col})

    possiblePaths := getPossiblePaths(row, col, pipes)

    for _, nextPoint := range possiblePaths {
        if !visited[nextPoint] {
            traverseMaze(nextPoint.x, nextPoint.y, pipes, visited, path)
        }
    }
}

func findCommonPosition(paths [][]Point) int {
    found := 0

    for i := 1; i < len(paths); i++ {
        for j:= 0; j < len(paths[i]); j++ {
            if paths[i - 1][j] == paths[i][j] {
                found = j + 1
                continue
            }
        }
    }
    return found
}

func getPossiblePaths(row, col int, pipes[][]string) []Point {
    start := []Point{}

    for _, p := range directions {
        nRow := row + p.x
        nCol := col + p.y

        if nRow < 0 || nRow > len(pipes) - 1 || nCol < 0 || nCol > len(pipes) - 1 {
            continue
        }

        if pipes[nRow][nCol] == "." {
            continue
        }

        next := pipes[nRow][nCol]
        curr := pipes[row][col]

        if exists := valid[p][curr][next]; exists{
            start = append(start, Point{nRow,nCol})
        }

    }

    return start
}

func parseTubes(tubes string) (int, int, [][]string) {

    splitPipes := strings.Split(tubes, "\n")
    var pipes [][]string
    var row, col int

    row = 0
    col = -1

    for _,sp := range splitPipes {
        if sp == "" {
            continue
        }

        split := strings.Split(sp, "")

        if col == -1 {
            col = strings.Index(sp, "S")
            row += 1
        }

        pipes = append(pipes, split)
    }

    return row - 1, col, pipes 
}
