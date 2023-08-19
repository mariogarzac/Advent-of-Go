package utils

import(
    "os"
    "log"
)

func OpenFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return file, nil
}

func min(sec []int) int {
    minNum := sec[0]

    for i := 1; i < len(sec); i++ {
        if sec[i] < minNum {
            minNum = sec[i]
        }
    }
    return minNum
}
