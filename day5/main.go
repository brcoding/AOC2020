package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
	"sort"
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
	var top, bottom, rowNumber, seatNumber, highestId float64
	seatRows := []float64{}
	for _, row := range input {
		bottom = 0
		top = 127

		for _, ru := range row[:7] {
			switch ru {
			case 'F':
				top = math.Floor(top / 2 + bottom / 2)
			case 'B':
				bottom = math.Round(bottom + ((top - bottom) / 2))
			}
		}
		rowNumber = top
		

		bottom = 0
		top = 7
		for _, ru := range row[7:] {
			switch ru {
			case 'L':
				top = math.Floor(top / 2 + bottom / 2)
			case 'R':
				bottom = math.Round(bottom + ((top - bottom) / 2))
			}
		}
		seatNumber = top
		seatId := rowNumber * 8 + seatNumber
		if seatId > highestId {
			highestId = seatId
		}
		seatRows = append(seatRows, seatId)
		// fmt.Printf("Row Number: %d\nSeat Number: %d\nSeat Id: %d\n", int(rowNumber), int(seatNumber), int(seatId))
	}
	fmt.Printf("Highest Seat Id: %d\n", int(highestId))
	sort.Float64s(seatRows)
	currRow := seatRows[0]
	for _, seat := range seatRows {
		if currRow != seat {
			fmt.Printf("Your Seat: %d\n", int(currRow))
			break
		}
		currRow++
	}
}
