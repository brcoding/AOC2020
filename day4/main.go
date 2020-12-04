package main

import (
	"fmt"
	"os"
	"bufio"
	"reflect"
	"strings"
	"sync"

	"gopkg.in/validator.v2"
)



type Passport struct {
	Byr string `validate:"nonzero"`
	Iyr string `validate:"nonzero"`
	Eyr string `validate:"nonzero"`
	Hgt string `validate:"nonzero"`
	Hcl string `validate:"nonzero"`
	Ecl string `validate:"nonzero"`
	Pid string `validate:"nonzero"`
	Cid string
}

func parsePassports(input []string, c chan<- Passport) {
	var passport Passport
	for _, line := range input {
		if line == "" {
			// Break for next passport
			fmt.Println(passport)
			c <- passport
			passport = Passport{}
			continue
		}
		for _, kvPair := range strings.Split(line, " ") {
			parts := strings.Split(kvPair, ":")
			fmt.Println(parts)
			// All of the below does NO type checking, do not use for anything important. Panic is inevitable.
			r := reflect.ValueOf(&passport).Elem()
    		f := r.FieldByName(strings.Title(strings.ToLower(parts[0])))
			f.SetString(parts[1])
		}
	}
	fmt.Println(passport)
	c <- passport
	passport = Passport{}
}

func (pass Passport) IsValid() (bool) {
	if errs := validator.Validate(pass); errs != nil {
		return false
	}
	return true
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

    wg := new(sync.WaitGroup)
	ch := make(chan Passport)

    wg.Add(1)
	go func(input []string) {
		defer wg.Done()
		parsePassports(input, ch)
	}(input)

	go func() {
		wg.Wait()
		close(ch)
	}()

	validCount := 0
    for passport := range ch {
    	if passport.IsValid() {
    		validCount++
    	}
        fmt.Println(passport.IsValid)
    }	

	fmt.Printf("Valid Passports: %d\n", validCount)
}

