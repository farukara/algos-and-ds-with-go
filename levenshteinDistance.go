package main

import (
    "fmt"
    "testing"
)

func main () {
    s1:="how are you"
    s2:="how old are you?"
    fmt.Println(Levenshtein(s1, s2))
}

func Levenshtein(s1, s2 string) int {
    matrix := make([][]int, len(s2)+1)
    //top left corner = 0
    matrix[0] = append(matrix[0], 0)

    //topmost row and leftmost column increment by 1
    for i:=1; i<len(s1)+1; i++ {
        matrix[0] = append(matrix[0], i)
    }
    for j:=1;j<len(s2)+1;j++ {
        matrix[j] = append(matrix[j], j)
    }

    //meat of the algorithm
    for i:=1; i<len(s1)+1; i++ {
        for j:=1; j<len(s2)+1; j++ {
            if s1[i-1] == s2[j-1] {
                //if characters are the same copy the top left corner = no work
                matrix[j] = append(matrix[j], matrix[j-1][i-1])
            } else {
                //if characters are not the same then take the minimum of preceding neighbours and add 1
                smallest := matrix[j][i-1]
                if matrix[j-1][i-1] < smallest { smallest = matrix[j-1][i-1]}
                if matrix[j-1][i] < smallest {smallest = matrix[j-1][i]}
                matrix[j] = append(matrix[j], smallest+1)
            }
        }
    }

    return matrix[len(s2)][len(s1)]
}

func test_Levenshtein(t *testing.T) {
    s1 := "saturday"
    s2 :="sunday"
    var distance int
    distance = Levenshtein(s1, s2)
    if distance != 3 {
        t.Errorf("(saturday, sunday) = %d; want 3", distance)
    }
    s1 = "abc"
    s2 ="bcd"
    distance = Levenshtein(s1, s2)
    if distance != 3 {
        t.Errorf("(saturday, sunday) = %d; want 3", distance)
    }
}

/* func test_Levenshtein_old () {
    s1 := "saturday"
    s2 :="sunday"
    var distance int
    distance = Levenshtein(s1, s2)
    if distance == 3 {
        fmt.Println("Distance is:", distance, ". Test PASSED")
    } else {
        fmt.Println("Returned:", distance, ", expecting: 3. Test !!!FAILED!!!.")
    }
    
    fmt.Println()
    s1 = "abc"
    s2 ="bcd"
    distance = Levenshtein(s1, s2)
    if distance == 2 {
        fmt.Println("Distance is:", distance, ". Test PASSED")
    } else {
        fmt.Println("Returned:", distance, ", expecting: 3. Test !!!FAILED!!!.")
    }
    func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
} */
