package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Day3() {
	data, err := ioutil.ReadFile("day3input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputs := strings.Split(string(data), "\n")

	count := oneTour(inputs, false, 3)

	fmt.Printf("Day 3, trees crashed into: %d\n", count)
}

func oneTour(inputs []string, skipOdd bool, offsetMultiplier int) int {
	count := 0
	for row, input := range inputs {
		if row == 0 || (skipOdd && row%2 > 0) {
			continue
		}
		if hasTreeAtOffset(row, input, offsetMultiplier, skipOdd) {
			count++
		}
	}
	return count
}

func hasTreeAtOffset(row int, part string, offsetMultiplier int, skipOdd bool) bool {
	partSize := len(part)
	offset := 0
	if skipOdd {
		offset = (row / 2) * offsetMultiplier
	} else {
		offset = row * offsetMultiplier
	}

	canvas := part
	if offset >= partSize {
		for i := 1; i < (offset/partSize)+1; i++ {
			canvas += part
		}
	}

	//fmt.Printf("%d %s (%d=%s)\n", row, canvas, offset, string(canvas[offset]))
	itemAtOffset := string(canvas[offset])
	return strings.Compare(itemAtOffset, `#`) == 0
}

func Day3Extra() {
	data, err := ioutil.ReadFile("day3input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputs := strings.Split(string(data), "\n")

	result := oneTour(inputs, false, 1)
	total := result
	result = oneTour(inputs, false, 3)
	total *= result
	result = oneTour(inputs, false, 5)
	total *= result
	result = oneTour(inputs, false, 7)
	total *= result
	result = oneTour(inputs, true, 1)
	total *= result

	fmt.Printf("Day 3+, total: %d\n", total)
}
