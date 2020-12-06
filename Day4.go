package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day4Extra() {
	data, err := ioutil.ReadFile("day4input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputs := strings.Split(string(data), "\n")

	count := 0
	var pass = passport{}
	for _, input := range inputs {
		if strings.Compare(input, "") == 0 {
			if isValidPassport(pass) {
				count++
			}
			pass = passport{}
		} else {
			pass = storeToPassport(pass, input)
		}
	}

	fmt.Printf("Day 4, passports ok: %d\n", count)
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func isValidPassport(pass passport) bool {
	//
	byr, err := strconv.Atoi(pass.byr)
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	//
	iyr, err := strconv.Atoi(pass.iyr)
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	//
	eyr, err := strconv.Atoi(pass.eyr)
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	//
	if strings.HasSuffix(pass.hgt, `cm`) {
		hgtCm, err := strconv.Atoi(strings.ReplaceAll(pass.hgt, `cm`, ``))
		if err != nil || hgtCm < 150 || hgtCm > 193 {
			return false
		}
	} else if strings.HasSuffix(pass.hgt, `in`) {
		hgtIn, err := strconv.Atoi(strings.ReplaceAll(pass.hgt, `in`, ``))
		if err != nil || hgtIn < 59 || hgtIn > 76 {
			return false
		}
	} else {
		return false
	}

	//
	if !regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(pass.hcl) {
		return false
	}

	//
	if !regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString(pass.ecl) {
		return false
	}

	//
	if !regexp.MustCompile(`^[0-9]{9}$`).MatchString(pass.pid) {
		return false
	}

	return true
}

func storeToPassport(pass passport, line string) passport {

	for _, token := range strings.Split(line, " ") {
		key := strings.Split(token, ":")[0]
		value := strings.Split(token, ":")[1]
		if strings.Compare(key, "byr") == 0 {
			pass.byr = value
		} else if strings.Compare(key, "iyr") == 0 {
			pass.iyr = value
		} else if strings.Compare(key, "eyr") == 0 {
			pass.eyr = value
		} else if strings.Compare(key, "hgt") == 0 {
			pass.hgt = value
		} else if strings.Compare(key, "hcl") == 0 {
			pass.hcl = value
		} else if strings.Compare(key, "ecl") == 0 {
			pass.ecl = value
		} else if strings.Compare(key, "pid") == 0 {
			pass.pid = value
		} else if strings.Compare(key, "cid") == 0 {
			pass.cid = value
		} else {
			fmt.Printf("Unexpected token: '%s'", token)
			os.Exit(1)
		}
	}
	return pass
}
