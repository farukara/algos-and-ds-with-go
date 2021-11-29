package main

import (
    "fmt"
)

func main() {
    fmt.Println(fib(50))
}

func fib (n int) int {
    memo := make(map[int]int)
    return fibInner(n, memo)
}

func fibInner(n int, memo map[int]int) int {
    f, ok := memo[n]
    if ok {return memo[n]}
    if n<=2 {return 1}
    f = fibInner(n-1, memo) + fibInner(n-2, memo)
    memo[n] = f
    return f
}
