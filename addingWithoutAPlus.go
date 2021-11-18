package main
//Adding 2 POSITIVE numbers without using "+" operator

import (
    "fmt"
)

func AddingWithoutPlus (a, b uint) uint {
    carry := (a & b) << 1
    total := a^b
    var temp uint
    for carry > 0 {
        temp  = total&carry << 1
        total = total^carry
        carry = temp
    }
    return total
}

func main() {
    const a uint = 12
    const b uint = 500
    fmt.Println(AddingWithoutPlus(a, b))
}
