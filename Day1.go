package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Day1() {
	data, err := ioutil.ReadFile("day1input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	listOfNumbers := strings.Split(string(data), "\n")

	for _, e1 := range listOfNumbers {
		for _, e2 := range listOfNumbers {
			e1i, _ := strconv.Atoi(e1)
			e2i, _ := strconv.Atoi(e2)

			if e1i+e2i == 2020 {
				fmt.Printf("Day 1, result is: %d\n", e1i*e2i)
				return
			}
		}
	}
}

func Day1Extra() {
	data, err := ioutil.ReadFile("day1input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	listOfNumbers := strings.Split(string(data), "\n")

	for _, e1 := range listOfNumbers {
		for _, e2 := range listOfNumbers {
			for _, e3 := range listOfNumbers {
				e1i, _ := strconv.Atoi(e1)
				e2i, _ := strconv.Atoi(e2)
				e3i, _ := strconv.Atoi(e3)

				if e1i+e2i+e3i == 2020 {
					fmt.Printf("Day 1+, result is: %d\n", e1i*e2i*e3i)
					return
				}
			}
		}
	}
}
