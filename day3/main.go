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

	slopes := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var slopeResults []int

	for _, slope := range slopes {
		fmt.Println(slope)
		rightIdx := 0

		trees := 0
		for idx, row := range input {
			if idx == 0 {
				fmt.Println(row + " Skipped")
				continue
			}
			if idx % slope[1] > 0 {
				fmt.Println(row + " Skipped")
				continue
			}
			rightIdx += slope[0]
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
		slopeResults = append(slopeResults, trees)
		trees = 0
	}

	totalTrees, productTrees :=	0, 1
	for _, result := range slopeResults {
		totalTrees += result
		productTrees *= result
	}
	
	fmt.Printf("Trees Hit: %d\nProduct: %d\n", totalTrees, productTrees)
}
