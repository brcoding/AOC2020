package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
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

	part2 := true
	year := 2020
	for _, i := range input {
		for _, ix := range input {
			if part2 {
				for _, iy := range input {
					if i + ix + iy == year {
						fmt.Printf("Answer: %d\n", i*ix*iy)
						return
					}
				}
			} else {
				if i+ix == year {
					fmt.Printf("Answer: %d\n", i*ix)
					return
				}				
			}
		}
	}

}