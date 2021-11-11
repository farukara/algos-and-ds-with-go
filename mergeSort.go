package main

import (
    "fmt"
    "time"
    "math/rand"
)

func mergeSort(s []int) []int {
    if len(s) <2 {return s}
    middle := len(s) >>1
    L := mergeSort(s[:middle])
    R := mergeSort(s[middle:])
    i, j := 0,0
    sorted_s := make([]int,0)
    for {
        if L[i] < R[j] {
            sorted_s = append(sorted_s, L[i])
            if i == len(L)-1 {
                sorted_s = append(sorted_s, R[j:]...)
                break
            }
            i += 1
        } else {
            sorted_s = append(sorted_s, R[j])
            if j == len(R)-1 {
                sorted_s = append(sorted_s,  L[i:]...)
                break
            }
            j += 1
        }
    }
    return sorted_s
}

func main() {
    rand.Seed(time.Now().Unix())
    slice := rand.Perm(10_000_000)
    fmt.Println("before:\n", slice)
    fmt.Println("after:")
    start := time.Now()
    fmt.Println("result:", mergeSort(slice))
    fmt.Println("Time:", time.Since(start).Seconds())
}
