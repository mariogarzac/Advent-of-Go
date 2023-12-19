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
    var i, j int
    i, j, pipes = parseTubes(string(file))

    fmt.Println(part1(i, j, pipes))

}

func part1(i, j int, pipes [][]string) int {

    begin := getPossiblePaths(i, j, pipes)
    paths := [][]Point{}

    for _, b := range begin {
        visited := make(map[Point]bool)
        path := []Point{}
        traverseMaze(b.x, b.y, pipes, visited, &path)
        paths = append(paths, path)
    }

    return findCommonPosition(paths)
}

func traverseMaze(i, j int, pipes [][]string, visited map[Point]bool, path *[]Point) {
    visited[Point{i, j}] = true

    // fmt.Printf("Visited point: (%d, %d) - Pipe: %s\n", i, j, pipes[i][j])
    *path = append(*path, Point{i,j})

    possiblePaths := getPossiblePaths(i, j, pipes)

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

func getPossiblePaths(i, j int, pipes[][]string) []Point {
    start := []Point{}

    for _, p := range directions {
        row := i + p.x
        col := j + p.y

        if row < 0 || row > len(pipes) - 1 || col < 0 || col > len(pipes) - 1 {
            continue
        }

        if pipes[row][col] == "." {
            continue
        }

        next := pipes[row][col]
        curr := pipes[i][j]

        if exists := valid[p][curr][next]; exists{
            start = append(start, Point{row,col})
        }

    }

    return start
}

func parseTubes(tubes string) (int, int, [][]string) {

    splitPipes := strings.Split(tubes, "\n")
    var pipes [][]string
    var i, j int

    i = 0
    j = -1

    for _,sp := range splitPipes {
        if sp == "" {
            continue
        }

        split := strings.Split(sp, "")

        if j == -1 {
            j = strings.Index(sp, "S")
            i += 1
        }

        pipes = append(pipes, split)
    }

    return i - 1, j, pipes 
}
