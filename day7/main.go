package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "regexp"
)

type Bag struct {
    BagType string
    Contains map[string]Bag
}

func makeBags(input []string) (map[string]Bag) {
    bagMap := make(map[string]Bag)
    for _, line := range input {
        bag := strings.Split(line, " bags contain ")[0]
        bagMap[bag] = Bag{bag, make(map[string]Bag, 0)}
    }
    return bagMap
}

func splitContents(contents string) ([]string) {
    reg, err := regexp.Compile("[^a-zA-Z ,]")
    if err != nil {
        panic(err)
    }
    contents = strings.ReplaceAll(contents, "bags", "")
    contents = strings.ReplaceAll(contents, "bag", "")
    cleanedContents := reg.ReplaceAllString(contents, "")

    var bagParts []string
    for _, part := range strings.Split(cleanedContents, ", ") {
        bagParts = append(bagParts, strings.TrimSpace(part))
    }

    return bagParts
}

func fillBags(input []string, bags *map[string]Bag) {
    for _, line := range input {
        lineParts := strings.Split(line, " bags contain ")
        bag := lineParts[0]
        bagContents := splitContents(lineParts[1])
        for _, item := range bagContents {
            if item == "no other" {
                continue
            }
            itemBag := (*bags)[bag]
            itemBag.Contains[item] = (*bags)[item]
            (*bags)[bag] = itemBag
        }
    }
    return 
}

func canHaveGoldBags(containerBag Bag) (bool) {
    for _, bag := range containerBag.Contains {
        if bag.BagType == "shiny gold" {
            return true
        } else {
            if canHaveGoldBags(bag) {
                return true
            }
        }
    }
    return false
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

    bags := makeBags(input)
    fillBags(input, &bags)
    bagCount := 0
    for _, bag := range bags {
        if canHaveGoldBags(bag) {
            bagCount++
        }
    }
    fmt.Printf("Gold Bags: %d\n", bagCount)
}

