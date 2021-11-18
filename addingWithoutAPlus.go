package main
//Adding 2 numbers without using "+" operator

import (
    "fmt"
)

func AddingWithoutPlus (a, b int) int {
    carry := (a & b) << 1
    total := a^b
    for carry > 0 {
        total = total^carry
        carry = (total ^ carry) << 1
    }
    return total
}

func main() {
    a := 500
    b := 10
    fmt.Println(AddingWithoutPlus(a, b))
}
