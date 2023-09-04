package eight

import (
    "Advent/utils"
    "fmt"
    "strconv"
    "strings"
)

func Solve(filename string){
    trees := FindTrees(filename)

    //part 1
    fmt.Println(isVisible(trees))

    //part 2
    fmt.Println(scenicScore(trees))
}

func FindTrees(filename string) [][]int {

    sc, file, _ := utils.OpenFile(filename)
    var trees [][]int

    defer file.Close()

    for sc.Scan() {
        treeLine := strings.Split(sc.Text(),"")

        var row []int
        for _,tree := range treeLine{
            value,_ := strconv.Atoi(tree)
            row = append(row, value)
        }
        trees = append(trees,row)
    }

    return trees
}

func scenicScore(trees [][]int) int {
    finish := len(trees)
    score := 0
    maxScore := 0


    for i := 1; i < finish - 1; i++ {
        for j := 1; j < finish - 1; j++ {
            score = calculateScore(trees, i, j)

            if maxScore < score{
                maxScore = score
                score = 0
            }
        }
    }
    return maxScore
}

func calculateScore(trees [][]int, row, col int) int {

    score := 1
    left,right,up,down := 0,0,0,0

    //Left
    for i := col - 1; i >= 0 ; i-- {
        left += 1
        if trees[row][col] <= trees[row][i] {
            break
        }
    }

    //Right
    for i := col + 1; i < len(trees); i++ {
        right += 1
        if trees[row][col] <= trees[row][i] {
            break
        }
    }

    //Up
    for i := row - 1; i >= 0; i-- {
        up += 1
        if trees[row][col] <= trees[i][col] {
            break
        }
    }

    //Down
    for i := row + 1; i < len(trees); i++ {
        down += 1
        if trees[row][col] <= trees[i][col] {
            break
        }
    }

    // fmt.Printf("left %d right %d up %d down %d \n", left,right,up,down)
    score *= left * right * up * down
    return score

}

func isVisible(trees [][]int) int {
    finish := len(trees)
    visible := finish * 2 + ((finish - 2) * 2)

    for i := 1; i < finish - 1; i++ {
        for j := 1; j < finish - 1; j++ {
            if checkVisible(trees, i, j) {
                visible += 1
            }
        }
    }

    return visible 
}

func checkVisible(trees [][]int, row, col int) bool {

    left, right, up, down := true, true, true, true

    for i := 0; i < len(trees); i++ {
        if i == col { continue }

        if i < col && trees[row][col] <= trees[row][i] {
            left = false
        }else if i > col && trees[row][col] <= trees[row][i] {
            right = false
        }
    }

    if left || right { return true }

    for i := 0; i < len(trees); i++ {
        if i == row {
            continue
        }
        if i < row && trees[row][col] <= trees[i][col] {
            up = false
        }else if i > row && trees[row][col] <= trees[i][col] {
            down = false
        }
    }

    if up || down { return true }

    return false
}


