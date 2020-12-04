package main

import (
    "fmt"
    "os"
    "bufio"
    "reflect"
    "strings"
    "sync"
    "strconv"
    "errors"

    "gopkg.in/validator.v2"
)


// type Passport struct {
//     Byr int `validate:"min=1920,max=2002"`
//     Iyr int `validate:"min=2010,max=2020"`
//     Eyr int `validate:"min=2020,max=2030"`
//     Hgt string `validate:"checkheight"`
//     Hcl string `validate:"regexp=^#[0-9a-fA-F]{6}"`
//     Ecl string `validate:"regexp=^(amb|blu|gry|grn|hzl|oth)$"`
//     Pid string `validate:"regexp=^\d{9}$"`
//     Cid string
// }

type Passport struct {
    Byr int `validate:"min=1920,max=2002"`
    Iyr int `validate:"min=2010,max=2020"`
    Eyr int `validate:"min=2020,max=2030"`
    Hgt string `validate:"checkheight"`
    Hcl string `validate:"regexp=^#[0-9a-fA-F]{6}$"`
    Ecl string `validate:"regexp=^(amb|blu|brn|gry|grn|hzl|oth)$"`
    Pid string `validate:"regexp=^[0-9]{9}$"`
    Cid string
}

func checkHeight(v interface{}, param string) error {
    st := reflect.ValueOf(v)
    if st.Kind() != reflect.String {
        return errors.New("notZZ only validates strings")
    }

    if strings.Contains(st.String(), "in") {
        height, _ := strconv.Atoi(strings.Split(st.String(), "in")[0])
        if height < 59 || height > 76 {
            return errors.New("Bad Height in Inches")
        }
    } else if strings.Contains(st.String(), "cm") {
        height, _ := strconv.Atoi(strings.Split(st.String(), "cm")[0])
        if height < 150 || height > 193 {
            return errors.New("Bad Height in cm")
        }
    } else {
        return errors.New("Not a valid height")
    }
    
    return nil
}

func parsePassports(input []string, c chan<- Passport) {
    var passport Passport
    for _, line := range input {
        if strings.TrimSpace(line) == "" {
            // Break for next passport
            c <- passport
            passport = Passport{}
            continue
        }
        for _, kvPair := range strings.Split(line, " ") {
            parts := strings.Split(strings.TrimSpace(kvPair), ":")
            // All of the below does NO type checking, do not use for anything important. Panic is inevitable.
            r := reflect.ValueOf(&passport).Elem()
            f := r.FieldByName(strings.Title(strings.ToLower(parts[0])))
            if f.Kind() == reflect.Int {
                i, _ := strconv.ParseInt(parts[1], 10, 64)
                f.SetInt(i)
            } else {
                f.SetString(parts[1])
            }
        }
    }
    c <- passport
    passport = Passport{}
}

func (pass Passport) IsValid() (bool, error) {
    validator.SetValidationFunc("checkheight", checkHeight)
    if errs := validator.Validate(pass); errs != nil {
        
        fmt.Printf("\n%+v\n%s\n\n", pass, errs)
        return false, errs
    }
    return true, nil
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
        if valid, _ := passport.IsValid(); valid {
            validCount++
        }
    }   

    fmt.Printf("Valid Passports: %d\n", validCount)
}

