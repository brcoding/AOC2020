package main

import (
	"fmt"
	"os"
	"bufio"
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

	rightIdx := 0

	trees := 0
	for idx, row := range input {
		if idx == 0 {
			continue
		}
		rightIdx += 3
		if len(row) <= rightIdx {
			rightIdx = rightIdx - len(row)
		}
		// tmpRow to print a pretty picture
		tmpRow := []byte(row)
		// # is a tree
		if string(row[rightIdx]) == "#" {
			tmpRow[rightIdx] = 'O'
			trees++
		}
		tmpRow[rightIdx] = 'X'
		fmt.Println(string(tmpRow))
	}
	fmt.Printf("Trees Hit: %d\n", trees)
}
