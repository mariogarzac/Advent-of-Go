package day01

import (
	"fmt"

    "github.com/mariogarzac/utils"
)

func FindNumbers() (int, error){
    sc, file, err := utils.OpenFile("day01/test.txt")
    if err != nil {
        return 0,err
    }

    var nums []byte
    count := 1

    for sc.Scan(){
        line := sc.Text()

        for i := range line{
            // ascii values for numbers go from 48 to 57
            if line[i] < 58{
                nums = append(nums, line[i] - 48)
            }
        }


        fmt.Println()

    }

    defer file.Close()

    return 0,nil
}

func main(){
    FindNumbers()
}
