package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Day6() {
	data, _ := ioutil.ReadFile("day6input.txt")
	inputs := strings.Split(string(data), "\n")

	totalAnswerCount := 0
	groupAnswerMap := make(map[int32]int)
	for _, lineOfAnswers := range inputs {
		if strings.Compare(lineOfAnswers, ``) == 0 {
			// group over
			totalAnswerCount += len(groupAnswerMap)
			groupAnswerMap = make(map[int32]int)
		} else {
			// append to group
			groupAnswerMap = parseAnswers(lineOfAnswers, groupAnswerMap)
		}
	}

	fmt.Printf("Day 6, total answers: %d\n", totalAnswerCount)
}

func parseAnswers(lineOfAnswers string, answerMap map[int32]int) map[int32]int {
	for _, y := range lineOfAnswers {
		answerMap[y]++
	}

	return answerMap
}

func Day6Extra() {
	data, _ := ioutil.ReadFile("day6input.txt")
	inputs := strings.Split(string(data), "\n")

	totalAnswerCount := 0
	membersPerGroup := 0
	groupAnswerMap := make(map[int32]int)
	for _, lineOfAnswers := range inputs {
		if strings.Compare(lineOfAnswers, ``) == 0 {
			// group over
			totalAnswerCount += countOfCommonAnswers(groupAnswerMap, membersPerGroup)
			groupAnswerMap = make(map[int32]int)
			membersPerGroup = 0
		} else {
			// append to group
			groupAnswerMap = parseAnswers(lineOfAnswers, groupAnswerMap)
			membersPerGroup++
		}
	}

	fmt.Printf("Day 6*, total common answers: %d\n", totalAnswerCount)
}

func countOfCommonAnswers(answerMap map[int32]int, memberCount int) int {

	count := 0
	for _, answerCount := range answerMap {
		if answerCount == memberCount {
			count++
		}
	}
	return count
}
