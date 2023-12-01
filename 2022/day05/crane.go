package five

import (
    "Advent/utils"
    "fmt"
    "log"
    "strconv"
    "strings"
)

type LoadingDock struct {
    Stacks [9]utils.Stack
}

func printStacks(ld *LoadingDock){
    for i := 0; i < len(ld.Stacks); i++ {
        fmt.Printf("%d ", i + 1)
        for j := 0; j < len(ld.Stacks[i]); j++ {
            fmt.Printf("%c ",ld.Stacks[i][j] )
        }
        fmt.Println()
    }
    fmt.Println("----------------------------")
}

/* 
This function takes the LoadingDock and line from file as parameters
It will then find the three values needed to move boxes to other stacks
amount is how many boxes are going to be moved, source is starting pile
and dest is destination pile
*/
func moveBoxes(ld *LoadingDock, text string) {

    textList := strings.Fields(text)

    var crane utils.Stack

    amount,err := strconv.Atoi(textList[1])
    if err != nil { log.Fatal(err) }

    source,err := strconv.Atoi(textList[3]) 
    if err != nil { log.Fatal(err) }

    dest,err := strconv.Atoi(textList[5]) 
    if err != nil { log.Fatal(err) }

    // Gather boxes
    for i := 0; i < amount; i++ {
        box, err := ld.Stacks[source - 1].Pop()
        if err != true { log.Fatal(err) }
        crane.Push(box)

    }

    // Insert boxes
    for i := 0; i < amount; i++ {
        box, err := crane.Pop()
        if err != true { log.Fatal(err) }

        ld.Stacks[dest - 1].Push(box)
    }


}

/*
This function reads the boxes from the input file character by character
and line by line. When it finds a letter, it will calculate its position to 
add it to the correct stack.
*/
func fillDock(ld *LoadingDock, text string){

    for i := 0; i < len(text); i++ {
        switch text[i]{
        case '[':
            // ignore
        case ']':
            // ignore
        case ' ':
            // ignore
        case '1','2','3','4','5','6','7','8','9':
            // ignore
        default:
            index := (i - 1)/4
            ld.Stacks[index].Insert(0,text[i])
        }
    }
}

/*
Returns the top of each stack 
*/
func getTops(ld *LoadingDock) utils.Stack {

    var tops utils.Stack
    for i := 0; i < len(ld.Stacks); i++ {
        tops.Push(ld.Stacks[i].Top())
    }

    return tops
}

/*
This function reads the file and passes each line to the correct function.
First it will fill the stacks with the boxes and then it will begin to move them.
Lastly it will return the top of each stack as the answer.
*/
func OperateCrane(filename string) utils.Stack{

    sc, file, err := utils.OpenFile(filename)
    if err != nil { log.Fatal(err) }

    defer file.Close()

    moveToInstructions := false

    ld := LoadingDock{}

    for sc.Scan(){

        if sc.Text() == "" {
            moveToInstructions = true
        }else if !moveToInstructions {
            fillDock(&ld, sc.Text())
        }else{
            moveBoxes(&ld, sc.Text())
        }
    }

    return getTops(&ld)
}
