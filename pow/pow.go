package main

import (
    "fmt"
    "os"
    "strconv"
)

func getLastDigit(n int) int {
    for n >= 10 {
        n %= 10
    }
    return n
}

func contains(s []int, x int) bool {
    for _, v := range s {
        if x == v {
            return true
        }
    }
    return false
}

func isOdd(n int) bool {
    return n % 2 != 0 && contains([]int{1, 3, 5, 7, 9}, getLastDigit(n))
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func nOddNumbers(n int) []int {
    n = abs(n)
    res := make([]int, n)
    for i, m := 0, 0; i < n; m++ {
        if isOdd(m) {
            res[i] = m
            i++
        }
    }
    return res
}

func pow(n int) int {
    res := 0
    for _, v := range nOddNumbers(n) {
        res += v
    }
    return res
}

func main() {
    arg, err := strconv.Atoi(os.Args[1])
    if err == nil {
        fmt.Print("- First n odd numbers: ")
        fmt.Println(nOddNumbers(arg))

        fmt.Print("- Power: ")
        fmt.Println(pow(arg))
    } else {
        fmt.Println(err)
    }
}