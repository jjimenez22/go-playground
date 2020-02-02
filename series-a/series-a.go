package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    arg, err := strconv.Atoi(os.Args[1])
    if err == nil {
        for i := 1; i <= arg; i++ {
            for j := 0; j < i; j++ {
                fmt.Printf("%d ", i)
            }
            fmt.Println()
        }
    } else {
        fmt.Println(err)
    }
}