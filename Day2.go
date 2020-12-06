package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const REGEX string = "(?P<least>\\d+)-(?P<most>\\d+) (?P<letter>[a-z]{1}): (?P<password>[a-z]+)"

func Day2() {
	data, err := ioutil.ReadFile("day2input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputs := strings.Split(string(data), "\n")

	okCount := 0
	for _, input := range inputs {
		atLeast, _ := strconv.Atoi(extractGroup(input, REGEX, "least"))
		atMost, _ := strconv.Atoi(extractGroup(input, REGEX, "most"))
		letter := extractGroup(input, REGEX, "letter")
		password := extractGroup(input, REGEX, "password")
		occurrences := strings.Count(password, letter)

		if occurrences >= atLeast && occurrences <= atMost {
			okCount++
		}
	}

	fmt.Printf("Day 2, result is %d\n", okCount)
}

func Day2Extra() {
	data, err := ioutil.ReadFile("day2input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputs := strings.Split(string(data), "\n")

	okCount := 0
	for _, input := range inputs {
		first, _ := strconv.Atoi(extractGroup(input, REGEX, "least"))
		second, _ := strconv.Atoi(extractGroup(input, REGEX, "most"))
		letter := extractGroup(input, REGEX, "letter")
		password := extractGroup(input, REGEX, "password")

		if len(input) == 0 {
			continue
		}

		a := ""
		if len(password) >= first {
			a = string(password[first-1])
		}
		b := ""
		if len(password) >= second {
			b = string(password[second-1])
		}
		matchA := strings.Compare(letter, a) == 0
		matchB := strings.Compare(letter, b) == 0
		if matchA != matchB {
			okCount++
		}
	}

	fmt.Printf("Day 2+, result is %d\n", okCount)
}

func extractGroup(data string, pattern string, groupName string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(data)

	for index, value := range r.SubexpNames() {
		if value == groupName && len(result) >= index {
			return result[index]
		}
	}
	return ``
}
