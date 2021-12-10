package main

import (
    "fmt"
)

func main() {
    weightLimit := 5
    weights := []int{5,3,4,2}
    values := []int{60,50,70,30}
    table := make([][]int, len(weights))
    for i:=0;i<len(weights);i++ {
        table[i] = append(table[i], 0)
    }
    for j:=1;j<=weightLimit;j++ {
        if j<weights[0] {
            table[0] = append(table[0], 0)
        } else {
            table[0] = append(table[0], values[0])
        }
    }
    for i:=1;i<len(weights);i++ {
        for j:=1;j<=weightLimit;j++ {
            if j<weights[i] {
                table[i] = append(table[i], table[i-1][j])
            } else {
                if table[i-1][j] > (values[i] + table[i-1][j-weights[i]]) {
                    table[i] = append(table[i], table[i-1][j])
                } else {
                    table[i] = append(table[i], (values[i] + table[i-1][j-weights[i]]))
                }
            }
        }
    }
    fmt.Println(table[len(weights)-1][weightLimit])
}
