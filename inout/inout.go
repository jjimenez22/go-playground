package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Print("-> ")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    fmt.Scanln("%s", &text)
    fmt.Println(text)
}