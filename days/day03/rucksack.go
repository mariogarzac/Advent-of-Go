package three

import (
	"Advent/utils"
	"fmt"
	"log"
)

func splitString(str string) (string, string)  {

    str1 := str[:len(str)/2]
    str2 := str[len(str)/2:]

    return str1, str2
}

func findCommonItem(s1 string, s2 string) byte {

    charset := make(map[byte]bool)

    for i := 0; i < len(s1); i++ {
        charset[s1[i]] = true
    }

    for i := 0; i < len(s2); i++ {
        if charset[s2[i]] {
            return s2[i]
        }
    }

    return 0
}

func RuckSack(filename string) int{

    sc, file, err := utils.OpenFile(filename)
    if err != nil { log.Fatal(err) }

    defer file.Close()

    priority := 0

    for sc.Scan(){
        text := sc.Text()

        // Split the string and find the common item
        s1,s2 := splitString(text)
        object := findCommonItem(s1, s2)

        priority += calculatePriority(object)
    }

    return priority
}

func RuckSackPart2(filename string) int {
    var contents []string
    counter, priority := 0, 0

    sc, file, err := utils.OpenFile(filename)
    if err != nil { log.Fatal(err) }

    defer file.Close()

    for sc.Scan(){
        if counter < 3{
            contents = append(contents, sc.Text())
            counter++
        }

        if counter == 3{
            counter = 0
            object := findCommonItemPart2(contents)
            priority += calculatePriority(object)

            contents = []string{}
        }
    }

    return priority
}

func findCommonItemPart2(c []string) byte {

    for i := 0; i < len(c[0]); i++ {
        for j := 0; j < len(c[1]); j++ {
            for k := 0; k < len(c[2]); k++ {
                if c[0][i] == c[1][j]  && c[1][j]  == c[2][k]{
                    return c[2][k]
                }
            }
        }
    }

    fmt.Println(c)
    return 0
}

func calculatePriority(char byte) int {
    if 'A' <= char && char <= 'Z'  {
        return int(char) - 65 + 27
    }

    if 'a' <= char && char <= 'z' {
        return int(char) - 97 + 1
    }

    return 0
}


