package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var (
		answerP1 int
		answerP2 int
	)
	year := 2020

	file, err := os.Open("input.txt")
 
	if err != nil {
		fmt.Printf("failed opening input file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input []int
 
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	for _, i := range input {
		for _, ix := range input {
			for _, iy := range input {
				if i + ix + iy == year {
					answerP2 = i * ix * iy
					break
				}
			}
			if i + ix == year {
				answerP1 = i * ix
				break
			}				
		}
	}
	fmt.Printf("Answer Part 1: %d\nAnswer Part 2: %d\n", answerP1, answerP2)
}