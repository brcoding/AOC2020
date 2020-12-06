package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func getMatchingChars(s1 string, s2 string) (string) {
    output := ""
    for _, s := range s1 {
        if strings.Contains(s2, string(s)) {
            output += string(s)
        }
    }
    return output
}

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

    // input = []string{"abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a", "a", "a", "", "b"}
    uniqueVals := make(map[rune]struct{})
    groupLetters := ""
    matchingGroupLetters := ""
    groupTotal := 0
    matchingGroupTotal := 0
    first := true
    for _, line := range input {
        if first {
            matchingGroupLetters = line
            first = false
        } else if strings.TrimSpace(line) != "" {
            matchingGroupLetters = getMatchingChars(matchingGroupLetters, line)        
        }
        groupLetters += line
        if strings.TrimSpace(line) == "" {
            for _, ru := range groupLetters {
                uniqueVals[ru] = struct{}{}
            }
            groupTotal += len(uniqueVals)
            matchingGroupTotal += len(matchingGroupLetters)

            uniqueVals = make(map[rune]struct{})
            matchingGroupLetters = ""
            first = true
            groupLetters = ""
        }
    }
    for _, ru := range groupLetters {
        uniqueVals[ru] = struct{}{}
    }
    groupTotal += len(uniqueVals)
    matchingGroupTotal += len(matchingGroupLetters)

    fmt.Printf("Group Total (Part 1): %d\nMatching Group Total (Part 2): %d\n", groupTotal, matchingGroupTotal)
}

