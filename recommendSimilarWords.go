package main

import (
    "fmt"
    "os"
    "log"
    "strings"
)

func main() {
    searchWord := "comple"
    fmt.Println(RecommendWords(searchWord))
}

func RecommendWords(searchWord string)[]string {
    //recommends 10 similar words to searchWord using LevenshteinDistnce algo
    file, err := os.ReadFile("p042_words.txt")
    if err != nil {
        log.Println("Error during opening the words file")
    }
    words := strings.Split(string(file)[1:len(file)-1], `","`)

    var scoreBoard = make(map[int][]string)
    score:=0
    for _, word := range words {
        if len(word) > len(searchWord) {
            score = Levenshtein(strings.ToLower(searchWord), strings.ToLower(word[:len(searchWord)]))
            scoreBoard[score] = append (scoreBoard[score], word)
        } else {
            score = Levenshtein(strings.ToLower(searchWord), strings.ToLower(word))
            scoreBoard[score] = append (scoreBoard[score], word)
        }
    }
    min := 10000
    for k, _ := range scoreBoard {
        if k<min {
            min= k
        }
    }
    var similarWords []string
    if len(scoreBoard[min]) > 10 {
        similarWords = append(similarWords, scoreBoard[min][:10]...)
        return similarWords
    } else {
        similarWords = append(similarWords, scoreBoard[min]...)
    }

    if len(scoreBoard[min+1]) > 10-len(similarWords) {
        similarWords = append(similarWords, scoreBoard[min+1][:10-len(similarWords)]...)
        return similarWords
    } else {
        similarWords = append(similarWords, scoreBoard[min+1]...)
    }
    
    if len(scoreBoard[min+2]) > 10-len(similarWords) {
        similarWords = append(similarWords, scoreBoard[min+2][:10-len(similarWords)]...)
        return similarWords
    } else {
        similarWords = append(similarWords, scoreBoard[min+2]...)
    }
    return similarWords
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
