package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Day5() {
	data, err := ioutil.ReadFile("day5input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputs := strings.Split(string(data), "\n")

	highestSeatId := 0
	for boardingPassIndex, boardingPass := range inputs {
		_, _, seatId := calcSeatIdForBoardingPass(boardingPass, boardingPassIndex)

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	//fmt.Printf("Day 5, highestSeatId: %d\n", highestSeatId)
}

func Day5Extra() {

	seatsByRow := make(map[int][]int)

	data, _ := ioutil.ReadFile("day5input.txt")
	inputs := strings.Split(string(data), "\n")

	for boardingPassIndex, boardingPass := range inputs {
		rowNumber, seatNumber, _ := calcSeatIdForBoardingPass(boardingPass, boardingPassIndex)

		seatsByRow[rowNumber] = append(seatsByRow[rowNumber], seatNumber)
	}

	for i := 1; i < 127; i++ {
		fmt.Printf("%d - %v\n", i, seatsByRow[i])
	}

}

func calcSeatIdForBoardingPass(boardingPass string, boardingPassIndex int) (int, int, int) {
	rowPart := boardingPass[:7]
	seatPart := boardingPass[7:]
	//fmt.Printf("Boarding pass %d: %s - %s\n", boardingPassIndex, rowPart, seatPart)

	rowNumber := binarySearch(0, 127, rowPart, `F`)
	seatNumber := binarySearch(0, 7, seatPart, `L`)
	seatId := (rowNumber * 8) + seatNumber
	//fmt.Printf("Pass #%d is located at row %d, seat %d (#id %d)\n", boardingPassIndex, rowNumber, seatNumber, seatId)
	return rowNumber, seatNumber, seatId
}

func binarySearch(front, back int, part, frontIndicator string) int {
	currentIndicator := string(part[0])
	middle := front + (back-front)/2
	//fmt.Printf("dig %s %d-%d-%d, looking at %s\n", part, front, middle, back, currentIndicator)
	if back-front <= 2 {
		if strings.Compare(currentIndicator, frontIndicator) == 0 {
			return front
		} else {
			return back
		}
	} else {
		if strings.Compare(currentIndicator, frontIndicator) == 0 {
			return binarySearch(front, middle, part[1:], frontIndicator)
		} else {
			return binarySearch(middle+1, back, part[1:], frontIndicator)
		}
	}
}
