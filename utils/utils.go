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


