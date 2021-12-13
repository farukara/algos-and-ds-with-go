package main

import (
    "fmt"
)

func Permutator (n int) [][]int {
//generates permutations of 1 thorough n digits
    var numbers []int
    var arrows []int //0 left, 1 rigth
    var movable []int
    var perms [][]int
    var dummy []int

    //create numbers
    for i:=1;i<=n;i++ {
        numbers = append(numbers, i)
        arrows = append(arrows, 0)
    }

    for {
        movable = nil
        //find the movable
        for i,number := range numbers {
            if arrows[i] == 0 {
                if i>0 && numbers[i-1] < number { 
                    movable = append(movable, number)
                }
            } else {
                if i<n-1 && numbers[i+1] < number {
                    movable = append(movable, number)
                }
            }
        }

        //append perms to the slice
        dummy = nil
        for _,number := range numbers {
            dummy = append(dummy, number)
        }
        perms = append(perms, dummy)
        if len(movable) == 0 {return perms}

        //find max movable number
        maxn := 0
        for _,number := range movable {
            if number > maxn {
                maxn = number
            }
        }

        //swap numbers and arrows
        for i := range numbers {
            if numbers[i] == maxn {
                if arrows[i] ==0 {
                    numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
                    arrows[i], arrows[i-1] = arrows[i-1], arrows[i]
                    break
                } else {
                    numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
                    arrows[i], arrows[i+1] = arrows[i+1], arrows[i]
                    break
                }
            }
        }
        //flip arrows direction
        for i,number := range numbers {
            if number>maxn {
                arrows[i] ^= 1
            }
        }
    }
}

func main() {
    fmt.Println(Permutator(9))
}
