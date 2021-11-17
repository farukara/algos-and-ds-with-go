package main

import (
    "fmt"
    "os"
    "strings"
)

//TODO: it's not differentiating between a longer word and the word itself.

func SimilarityScore(s1, s2 string) int {
    score := 0
    matrix := make([][]int, len(s1))
    for i := range matrix {
        matrix[i] = make([]int, len(s2))
    }
    for i:=0; i<len(s1); i++{
        for j:=0; j<len(s2); j++{
            if s1[i] == s2[j] {
                if i >1 && j>1 {
                    matrix[i][j] = matrix[i-1][j-1] +1
                    if matrix[i][j] > score {
                        score = matrix[i][j]
                    }
                } else {
                    matrix[i][j] = 1
                    if matrix[i][j] > score {
                        score = matrix[i][j]
                    }
                }
            } else {
                matrix[i][j] = 0
            }
        }
    }
    return score
}

func find_closest_match(s1 string) string  {
    //read-in the words 
    words_file, err := os.ReadFile("p042_words.txt")
    if err != nil {
        fmt.Println("Error during opening words file")
    }
    words := strings.Split(string(words_file)[1:len(words_file)-1], `","`)

    //find the closest match
    score := 0
    target_word := ""
    for _, word := range words {
        how_similar := SimilarityScore(strings.ToLower(s1), strings.ToLower(word))
        if how_similar > score {
            fmt.Println(word)
            score = how_similar
            target_word = word
        }
    }
    return target_word
}

func main() {
    s1 := "establish"
    /* s2 := "fish"
    s3 := "vista"
    s4 := "tsracadabra"
    fmt.Println(SimilarityScore(s1,s2))
    fmt.Println(SimilarityScore(s1,s3))
    fmt.Println(SimilarityScore(s1,s4)) */
    fmt.Println("Closest match is: ", find_closest_match(s1))
}
