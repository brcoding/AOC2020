package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    file, err := os.Open("input.txt")
 
    if err != nil {
        fmt.Printf("failed opening input file: %s", err)
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var input []string
 
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    uniqueVals := make(map[rune]struct{})
    groupLetters := ""
    groupTotal := 0
    for _, line := range input {
        groupLetters += line
        if strings.TrimSpace(line) == "" {
            for _, ru := range groupLetters {
                uniqueVals[ru] = struct{}{}
            }
            groupTotal += len(uniqueVals)
            fmt.Printf("L: %s\nG: %d\n", groupLetters, groupTotal)
            uniqueVals = make(map[rune]struct{})
            groupLetters = ""
        }
    }
    for _, ru := range groupLetters {
        uniqueVals[ru] = struct{}{}
    }
    groupTotal += len(uniqueVals)

    fmt.Println(groupTotal)
}

