package main

import (
    "testing"
    "errors"
)

func isErr(err error, val string) (bool) {
    if err == nil && val == "" {
        return false
    }
    if err.Error() != errors.New(val).Error() {
        return true
    }
    return false
}

func TestPassportValidation(t *testing.T) {
    // Byr
    passport := Passport{Byr:1901, Iyr:2014, Eyr:2027, Hgt:"183cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Byr: less than min") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:3101, Iyr:2014, Eyr:2027, Hgt:"183cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Byr: greater than max") {
        t.Errorf("Failed: %s", err)
    }
    
    // Iyr
    passport = Passport{Byr:1980, Iyr:2024, Eyr:2027, Hgt:"183cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Iyr: greater than max") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:1024, Eyr:2027, Hgt:"183cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Iyr: less than min") {
        t.Errorf("Failed: %s", err)
    }

    // Eyr
    passport = Passport{Byr:1980, Iyr:2020, Eyr:1990, Hgt:"183cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Eyr: less than min") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2035, Hgt:"183cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Eyr: greater than max") {
        t.Errorf("Failed: %s", err)
    }

    // height
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"183",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hgt: Not a valid height") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hgt: Not a valid height") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"cm123",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hgt: Bad Height in cm") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"29cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hgt: Bad Height in cm") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"in123",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hgt: Bad Height in Inches") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"123in",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hgt: Bad Height in Inches") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"59in",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"76in",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"150cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"193cm",
        Hcl:"#aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "") {
        t.Errorf("Failed: %s", err)
    }

    // Hcl
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"#aaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hcl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    }    
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"#333", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hcl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"#aaaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hcl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hcl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"##aaaaaa", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hcl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    }
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"", Ecl:"grn", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Hcl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    } 

    // Ecl
    passport = Passport{Byr:1980, Iyr:2020, Eyr:2020, Hgt:"155cm",
        Hcl:"#aaaaaa", Ecl:"asdf", Pid:"974469723", Cid:"176"}
    if _, err := passport.IsValid(); isErr(err, "Ecl: regular expression mismatch") {
        t.Errorf("Failed: %s", err)
    } 


}