package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type PasswordParts struct {
	MinOccurrence  int
	MaxOccurrence  int
	Letter		   string
	Password 	   string
}

func getPasswordParts(pwinput string) (PasswordParts) {
	firstParts := strings.Split(pwinput, "-")
	min, err := strconv.Atoi(firstParts[0])
	if err != nil {
		fmt.Println(pwinput)
		panic(err)
	}
	secondParts := strings.Split(strings.Join(firstParts[1:], ""), " ")
	max, err := strconv.Atoi(secondParts[0])
	if err != nil {
		fmt.Println(pwinput)
		panic(err)
	}
	thirdParts := strings.Split(strings.Join(secondParts[1:], " "), ": ")
	letter := thirdParts[0]
	if err != nil {
		fmt.Println(pwinput)
		panic(err)
	}
	password := strings.Join(thirdParts[1:], ": ")

	return PasswordParts{min, max, letter, password}
}

func matches(letter string, password string, pos int) (bool) {
	if string(password[pos - 1]) == letter {
		return true
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

	validPasswords := 0
	validPasswordsPart2 := 0
	for _, pwinput := range input {
		parts := getPasswordParts(pwinput)
		letterCount := strings.Count(parts.Password, parts.Letter)
		if letterCount >= parts.MinOccurrence && letterCount <= parts.MaxOccurrence {
			validPasswords++
		}
		minMatch := matches(parts.Letter, parts.Password, parts.MinOccurrence)
		maxMatch := matches(parts.Letter, parts.Password, parts.MaxOccurrence)
		if (minMatch != maxMatch) {
			validPasswordsPart2++
		}
	}
	fmt.Printf("Valid Passwords: %d\nValid Passwords Part 2: %d\n", validPasswords, validPasswordsPart2)
}
