package main

import (
    "fmt"
)

func main() {
    fmt.Println(fib(50))
    
}

func fib(n int) int {
    memo := make(map[int]int)
    memo[1] = 1
    memo[2] = 1

    if n<=2 {
        return 1
    } else {
        for i:=3; i<=n; i++ {
            memo[i] = memo[i-1] + memo[i-2]
        }
    }
    return memo[n]
}
