package main

import (
    "fmt"
    "math/rand"
    "time"
)

func InsertionSort(slice []int) []int {
    for j:=1; j<len(slice); j++ {
        for i:=j; i>0; i--{
            if slice[i] < slice[i-1] {
                slice[i], slice[i-1] = slice[i-1], slice[i]
        } else {
            break
        }
        }
    }
    return slice
}

func test_InsertionSort() {
    test_data := make([][]int,0)
    test_data = append(test_data, []int{5,3,1,4,2})
    test_data = append(test_data, []int{})
    test_data = append(test_data, []int{-1,-2,0,1,3})

    rand.Seed(time.Now().Unix())
    test_data = append(test_data, rand.Perm(100))

    start := time.Now()
    for _, slice := range test_data {
        fmt.Println("Unsorted:", slice)
        fmt.Println("Sorted:", InsertionSort(slice), "\n")
    }
    fmt.Println("Time:", time.Since(start).Seconds())
}

func main () {
    test_InsertionSort()
}

